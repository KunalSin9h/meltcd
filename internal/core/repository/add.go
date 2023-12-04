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
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/charmbracelet/log"
)

type Repository struct {
	URL, Secret string
}

func (r *Repository) saveCredential(username, password string) {
	r.Secret = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func (r *Repository) getCredential() (username, password string) {
	d, err := base64.StdEncoding.DecodeString(r.Secret)
	if err != nil {
		log.Error(err.Error())
		return "", ""
	}

	cred := strings.Split(string(d), ":")
	if len(cred) != 2 {
		log.Error("repository credential is empty")
		return "", ""
	}

	username = cred[0]
	password = cred[1]

	return username, password
}

var repositories []*Repository

func Add(url, username, password string) error {
	repo, found := findRepo(url)
	if found {
		return errors.New("repository with same url already exists")
	}

	repo.saveCredential(username, password)

	repositories = append(repositories, &Repository{
		URL:    url,
		Secret: repo.Secret,
	})
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
	var repos []*Repository

	if err := json.Unmarshal(*d, &repos); err != nil {
		return err
	}

	repositories = repos
	return nil
}
