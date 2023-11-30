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

	"github.com/meltred/meltcd/internal/core/secrets"
)

type Repository struct {
	URL, SecretId string
}

var repositories []*Repository

func Add(url, username, password string) error {
	id, err := secrets.CreateRepository(url, username, password)
	if err != nil {
		return err
	}

	repositories = append(repositories, &Repository{
		URL:      url,
		SecretId: id,
	})
	return nil
}

func GetRepoData() ([]byte, error) {
	result, err := json.Marshal(repositories)
	if err != nil {
		return []byte{}, err
	}

	return result, nil
}

func LoadRepoData(d *[]byte) error {
	var repos []*Repository

	if err := json.Unmarshal(*d, &repos); err != nil {
		return err
	}

	repositories = repos
	return nil
}
