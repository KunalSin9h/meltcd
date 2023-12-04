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
	"strings"

	"github.com/charmbracelet/log"
)

func Find(repoURL string) (string, string) {
	repoURL, _ = strings.CutSuffix(repoURL, "/")

	repo, found := findRepo(repoURL)
	if !found {
		return "", ""
	}

	username, password := repo.getCredential()

	if username == "" || password == "" {
		log.Error("username and password not found in secret")
		return "", ""
	}

	return username, password
}

func findRepo(url string) (*Repository, bool) {
	for _, x := range repositories {
		if x.URL == url || x.URL+".git" == url || x.URL == url+".git" {
			return x, true
		}
	}

	return &Repository{}, false
}
