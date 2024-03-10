/*
Copyright 2023 - PRESENT Meltred

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"log/slog"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

type Repository struct {
	URL, ImageRef, Secret string
	Reachable             bool
}

var repositories []*Repository

func (r *Repository) saveCredential(username, password string) {
	r.Secret = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func (r *Repository) GetRegistryAuth() (string, error) {
	username, password := r.getCredential()

	payload := map[string]string{
		"username": username,
		"password": password,
	}

	d, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(d), nil
}

func (r *Repository) getCredential() (username, password string) {
	d, err := base64.StdEncoding.DecodeString(r.Secret)
	if err != nil {
		slog.Error(err.Error())
		return "", ""
	}

	cred := strings.Split(string(d), ":")
	if len(cred) != 2 {
		slog.Error("repository credential is empty")
		return "", ""
	}

	username = cred[0]
	password = cred[1]

	return username, password
}

func (r *Repository) checkReachability(username, password string) {
	if r.URL != "" {
		fs := memfs.New()
		storage := memory.NewStorage()

		_, err := git.Clone(storage, fs, &git.CloneOptions{
			URL:          r.URL,
			SingleBranch: true,
			Depth:        1,
			Auth: &http.BasicAuth{
				Username: username,
				Password: password,
			},
		})

		if err != nil {
			r.Reachable = false
		}
	} else if r.ImageRef != "" {
		cli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			slog.Error(err.Error())
			r.Reachable = false
			return
		}

		auth, err := r.GetRegistryAuth()
		fmt.Println(auth)
		if err != nil {
			slog.Error(err.Error())
			r.Reachable = false
			return
		}

		_, err = cli.ImagePull(context.Background(), r.ImageRef, types.ImagePullOptions{
			RegistryAuth: auth,
		})

		if err != nil {
			slog.Error(err.Error())
			r.Reachable = false
			return
		}

		r.Reachable = true
	} else {
		slog.Error("Both url and image ref is empty")
	}
}

// url is git url or container image name
func Add(url, imageRef, username, password string) error {
	// since eight url is there of imageRef,
	name := url + imageRef
	repo, found := FindRepo(name)
	if found {
		return errors.New("repository with same url already exists")
	}

	repo.URL = url
	repo.ImageRef = imageRef
	repo.saveCredential(username, password)
	repo.Reachable = true

	go repo.checkReachability(username, password)

	repositories = append(repositories, repo)
	return nil
}

func GetData() ([]byte, error) {
	result, err := json.Marshal(repositories)
	if err != nil {
		return []byte{}, err
	}

	return result, nil
}

func LoadData(d *[]byte) error {
	return json.Unmarshal(*d, &repositories)
}
