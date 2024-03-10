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
	"log/slog"
)

func FindCreds(repoName string) (string, string) {
	repo, found := FindRepo(repoName)
	if !found {
		return "", ""
	}

	username, password := repo.getCredential()

	if username == "" || password == "" {
		slog.Error("username and password not found in secret")
		return "", ""
	}

	return username, password
}

func FindRepo(name string) (*Repository, bool) {
	for _, x := range repositories {
		if x.URL == name || x.URL+".git" == name || x.URL == name+".git" || x.ImageRef == name {
			return x, true
		}
	}

	return &Repository{}, false
}
