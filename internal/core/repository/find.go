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
	"strings"

	"github.com/charmbracelet/log"
)

func Find(repoURL string) (string, string) {
	repoURL, _ = strings.CutSuffix(repoURL, "/")

	secret, found := findSecret(repoURL)
	if !found {
		return "", ""
	}

	d, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		log.Error(err.Error())
		return "", ""
	}

	cred := strings.Split(string(d), ":")

	if len(cred) < 2 {
		log.Error("username and password not found in secret")
		return "", ""
	}

	return cred[0], cred[1]
}

func findSecret(url string) (string, bool) {
	for _, x := range repositories {
		if x.URL == url || x.URL+".git" == url || x.URL == url+".git" {
			return x.Secret, true
		}
	}

	return "", false
}
