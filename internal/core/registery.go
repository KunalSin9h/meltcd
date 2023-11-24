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

	_, exists := getApp(app.Name)
	if exists {
		return fmt.Errorf("app already exists with name: %s", app.Name)
	}

	go app.Run()
	Applications = append(Applications, app)

	log.Info("Registered!")
	return nil
}

func Update(app *application.Application) error {
	log.Info("Updating application", "name", app.Name)

	runningApp, exists := getApp(app.Name)
	if !exists {
		return fmt.Errorf("app does not exists, create a new application first")
	}

	runningApp.RefreshTimer = app.RefreshTimer
	runningApp.Source = app.Source

	log.Info("Updated!")
	return nil
}

func Details(app_name string) (application.Application, error) {
	log.Info("Getting application details", "name", app_name)

	runningApp, exists := getApp(app_name)
	if !exists {
		return application.Application{}, fmt.Errorf("app does not exists, create a new application first")
	}

	log.Info("Done!")
	return *runningApp, nil
}

func getApp(name string) (*application.Application, bool) {
	for _, app := range Applications {
		if app.Name == name {
			return app, true
		}
	}

	return &application.Application{}, false
}
