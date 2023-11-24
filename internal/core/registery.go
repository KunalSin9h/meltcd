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

package core

import (
	"fmt"
	"meltred/meltcd/internal/core/application"

	"github.com/charmbracelet/log"
)

var Applications []*application.Application

func Register(app *application.Application) error {
	log.Info("Registering application", "name", app.Name)

	for _, regApp := range Applications {
		if regApp.Name == app.Name {
			return fmt.Errorf("app already exists with name: %s", app.Name)
		}
	}

	go app.Run()
	Applications = append(Applications, app)

	return nil
}

func Update(app *application.Application) error {
	log.Info("Updating application", "name", app.Name)

	exists := false

	for _, regApp := range Applications {
		if regApp.Name == app.Name {

			regApp.RefreshTimer = app.RefreshTimer
			regApp.Source = app.Source

			exists = true
			break
		}
	}

	if !exists {
		return fmt.Errorf("app does not exists, create a new application first")
	}
	return nil
}
