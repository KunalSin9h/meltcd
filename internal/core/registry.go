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
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/meltred/meltcd/internal/core/application"

	"github.com/charmbracelet/log"
)

var Applications []*application.Application

func Register(app *application.Application) error {
	log.Info("Registering application", "name", app.Name)

	_, exists := getApp(app.Name)
	if exists {
		return fmt.Errorf("app already exists with name: %s", app.Name)
	}

	app.SyncTrigger = make(chan application.SyncType, 1)

	timeOfCreation := time.Now()
	app.CreatedAt = timeOfCreation
	app.UpdatedAt = timeOfCreation
	app.LastSyncedAt = timeOfCreation

	go app.Run()
	Applications = append(Applications, app)

	log.Info("Registered!")
	return nil
}

// TODO: update atomic parts
// only specify things you need go update
func Update(app *application.Application) error {
	log.Info("Updating application", "name", app.Name)

	runningApp, exists := getApp(app.Name)
	if !exists {
		return fmt.Errorf("app does not exists, create a new application first")
	}

	runningApp.RefreshTimer = app.RefreshTimer
	runningApp.Source = app.Source

	runningApp.UpdatedAt = time.Now()

	// Sync the application as new update is done
	runningApp.SyncTrigger <- application.UpdateSync

	return nil
}

func Details(appName string) (application.Application, error) {
	runningApp, exists := getApp(appName)
	if !exists {
		return application.Application{}, fmt.Errorf("app does not exists, create a new application first")
	}

	runningApp.HealthStatus = runningApp.Health.ToString()

	return *runningApp, nil
}

type AppList struct {
	Data []AppStatus `json:"data"`
}

type AppStatus struct {
	ID           uint32    `json:"id"`
	Name         string    `json:"name"`
	Health       string    `json:"health"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAT    time.Time `json:"updated_at"`
	LastSyncedAt time.Time `json:"last_synced_at"`
}

func List() AppList {
	var res AppList

	for index, app := range Applications {
		res.Data = append(res.Data, AppStatus{
			ID:           uint32(index),
			Name:         app.Name,
			Health:       app.Health.ToString(),
			CreatedAt:    app.CreatedAt,
			UpdatedAT:    app.UpdatedAt,
			LastSyncedAt: app.LastSyncedAt,
		})
	}

	return res
}

func getApp(name string) (*application.Application, bool) {
	for _, app := range Applications {
		if app.Name == name {
			return app, true
		}
	}

	return &application.Application{}, false
}

func Refresh(appName string) error {
	app, exists := getApp(appName)
	if !exists {
		return fmt.Errorf("app does not exists, create a new application first")
	}

	app.SyncTrigger <- application.Synchronize

	return nil
}

func getRegistryData() ([]byte, error) {
	result, err := json.Marshal(Applications)
	if err != nil {
		return []byte{}, err
	}

	return result, nil
}

func loadRegistryData(d *[]byte) error {
	var load []*application.Application

	if err := json.Unmarshal(*d, &load); err != nil {
		return err
	}

	for _, app := range load {
		app.SyncTrigger = make(chan application.SyncType, 1)
		go app.Run()
	}

	Applications = load

	return nil
}

func RemoveApplication(appName string) error {
	go removeSvcFromApps(appName)

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	runningService, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return err
	}

	// set (unique) of network to remove
	networksToRemove := map[string]bool{}

	for _, svc := range runningService {
		name := svc.Spec.Labels["com.docker.stack.namespace"]

		if name == appName {
			if err := cli.ServiceRemove(context.Background(), svc.ID); err != nil {
				return err
			}

			for _, nets := range svc.Spec.TaskTemplate.Networks {
				// nets.Target is the network ID
				networksToRemove[nets.Target] = true
			}
		}
	}

	for networkID := range networksToRemove {
		if err := cli.NetworkRemove(context.Background(), networkID); err != nil {
			return err
		}
	}

	return nil
}

func removeSvcFromApps(appName string) {
	tmp := make([]*application.Application, 0)

	for _, app := range Applications {
		if app.Name != appName {
			tmp = append(tmp, app)
		}
	}

	Applications = tmp
}

func Recreate(appName string) error {
	data, err := Details(appName)
	if err != nil {
		return err
	}
	log.Info("Got details of application", "app_name", appName)

	// clearing the current state, so it can be recreated
	data.LiveState = ""

	if err := RemoveApplication(appName); err != nil {
		return err
	}
	log.Info("Removed application", "app_name", appName)

	return Register(&data)
}
