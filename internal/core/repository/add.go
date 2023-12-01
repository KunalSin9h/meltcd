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
	"encoding/json"
	"net/url"

	"github.com/meltred/meltcd/internal/core/secrets"
)

type Repository struct {
	URL, SecretID string
}

var repositories []*Repository

func Add(uri, username, password string) error {
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}

	id, err := secrets.CreateRepository(u, username, password)
	if err != nil {
		return err
	}

	repositories = append(repositories, &Repository{
		URL:      u.String(),
		SecretID: id,
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
