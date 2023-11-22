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

package application

import (
	"errors"
)

type Spec struct {
	Name         string `json:"name" yaml:"name"`
	RefreshTimer string `json:"refresh_timer" yaml:"refresh_timer"` // number of minutes
	Source       Source `json:"source" yaml:"source"`
}

type Source struct {
	RepoURL        string `json:"repo" yaml:"repoURL"`
	TargetRevision string `json:"revision" yaml:"targetRevision"`
	Path           string `json:"path" yaml:"path"`
}

// parse an application from yaml source
func ParseSpecFromFile(file string) (Spec, error) {
	if file == "" {
		return Spec{}, errors.New("Application specification file not specified")
	}

	return Spec{}, nil
}

func ParseSpecFromValue(name, repo, revision, path, refresh string) (Spec, error) {
	if repo == "" {
		return Spec{}, errors.New("the git repository not specified")
	}

	if path == "" {
		return Spec{}, errors.New("the path to Service file not specified")
	}

	return Spec{
		Name:         name,
		RefreshTimer: refresh,
		Source: Source{
			RepoURL:        repo,
			TargetRevision: revision,
			Path:           path,
		},
	}, nil
}
