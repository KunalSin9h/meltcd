/*
Copyright 2023 - PRESENT kunalsin9h

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
	"bytes"
	"context"
	"errors"
	"strings"
	"time"

	"log/slog"

	"github.com/kunalsin9h/meltcd/internal/core/repository"
	"github.com/kunalsin9h/meltcd/spec"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"gopkg.in/yaml.v2"
)

type Application struct {
	ID           uint32        `json:"id"`
	Name         string        `json:"name"`
	Source       Source        `json:"source"`
	RefreshTimer string        `json:"refresh_timer"` // Timer to check for Sync format of "3m50s"
	Health       Health        `json:"health"`
	HealthStatus string        `json:"health_status"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	LastSyncedAt time.Time     `json:"last_synced_at"`
	LiveState    string        `json:"-"`
	SyncTrigger  chan SyncType `json:"-"`
}

type Health int

const (
	Healthy Health = iota
	Progressing
	Degraded
	Suspended
)

func (h Health) ToString() string {
	switch h {
	case Healthy:
		return "healthy"
	case Progressing:
		return "progressing"
	case Degraded:
		return "degraded"
	case Suspended:
		return "suspended"
	}

	return "NA"
}

type SyncType int

const (
	Synchronize SyncType = iota
	UpdateSync
)

func New(spec Spec) Application {
	return Application{
		Name:         spec.Name,
		RefreshTimer: spec.RefreshTimer,
		Source:       spec.Source,
	}
}

func (app *Application) Run() {
	slog.Info("Running Application", "name", app.Name)

	ticker := time.NewTicker(time.Minute * 3)
	defer ticker.Stop()

	// updateTicker is for updating the refresh timer, since app settings
	// can be changed at run time (Run() function) so we have to update the timer in every loop.
	if err := updateTicker(app.RefreshTimer, ticker); err != nil {
		slog.Error(err.Error())
		app.Health = Suspended
		return
	}

	slog.Info("Staring sync process")

	for ; true; waitSync(ticker.C, app.SyncTrigger) {
		if err := updateTicker(app.RefreshTimer, ticker); err != nil {
			slog.Error(err.Error())
			app.Health = Degraded
			continue
		}

		targetState, err := app.GetState()
		if err != nil {
			slog.Warn("Not able to get service", "repo", app.Source.RepoURL)
			slog.Error(err.Error())
			app.Health = Degraded
			continue
		}
		slog.Info("got target state")
		if app.SyncStatus(targetState) {
			// TODO: Sync Status = Synched
			slog.Info("Synched")
			app.Health = Healthy
			continue
		}
		slog.Info("liveState and Target state is out of sync. syncing now...")

		// // TODO: Sync Status = Out of Sync
		app.Health = Progressing
		if err := app.Apply(targetState); err != nil {
			app.Health = Degraded
			slog.Warn("Not able to apply targetState", "error", err.Error())
			continue
		}

		app.Health = Healthy
		slog.Info("Applied new changes")
	}
}

func waitSync(ticker <-chan time.Time, syncTrigger <-chan SyncType) {
	select {
	case <-ticker:
	case <-syncTrigger:
	}
}

func updateTicker(duration string, t *time.Ticker) error {
	refreshTime, err := time.ParseDuration(duration)
	if err != nil {
		slog.Error("Failed to parse refresh_time, it must be like \"3m30s\"", "name", duration)
		return err
	}

	t.Reset(refreshTime)
	return nil
}

func (app *Application) GetState() (string, error) {
	slog.Info("Getting service state from git repo", "repo", app.Source.RepoURL, "app_name", app.Name)
	// TODO: not using targetRevision

	// TODO: IMPROVEMENT
	// Use Docker Volumes to clone repository
	// and then only fetch & pull if already exists
	// and check if specified path is modified then apply the changes
	fs := memfs.New()
	storage := memory.NewStorage()
	// defer clear storage, i (kunal singh) think that when storage goes out-of-scope
	// it is cleared

	username, password := repository.FindCreds(app.Source.RepoURL)

	// TODO: Improvement
	// GET the name and commit also
	// so that we can show it in the ui or something
	ref := plumbing.HEAD
	if app.Source.TargetRevision != "HEAD" {
		ref = plumbing.NewBranchReferenceName(app.Source.TargetRevision)
	}

	_, err := git.Clone(storage, fs, &git.CloneOptions{
		URL:           app.Source.RepoURL,
		ReferenceName: ref,
		SingleBranch:  true,
		Depth:         1,
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
	})

	// if errors.Is(err, git.ErrRepositoryAlreadyExists) {
	// 	//  fetch & pull request
	// 	// don't clone again
	// 	slog.Info("Repo already exits", "repo", app.Source.RepoURL)
	// 	slog.Error("Since the storage is not persistent, this error should not exist")
	// } else
	if err != nil {
		return "", err
	}

	serviceFile, err := fs.Open(app.Source.Path)
	if err != nil {
		slog.Error("Path not found", "repo", app.Source.RepoURL, "path", app.Source.Path)
		return "", err
	}
	defer serviceFile.Close()

	// reading the service file content
	buf := new(bytes.Buffer)
	buf.ReadFrom(serviceFile)

	return buf.String(), nil
}

func (app *Application) Apply(targetState string) error {
	slog.Info("Applying new targetState")
	// TODO this client can be stored i app or new struct core
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		slog.Error("Not able to create a new docker client")
		return err
	}

	var swarmSpec spec.DockerSwarm
	if err := yaml.Unmarshal([]byte(targetState), &swarmSpec); err != nil {
		return err
	}

	// TODO use volOpts
	// create volume
	for volName, volOpts := range swarmSpec.Volumes {
		labels := make(map[string]string)
		for _, l := range volOpts.Labels {
			tokens := strings.SplitN(l, "=", 2)
			if len(tokens) != 2 {
				return errors.New("invalid labels in volume")
			}

			labels[tokens[0]] = tokens[1]
		}

		cli.VolumeCreate(context.Background(), volume.CreateOptions{
			Name:       volName,
			Driver:     volOpts.Driver,
			DriverOpts: volOpts.DriverOpts,
			Labels:     labels, // TODO labels are not working
		})
	}

	networkID, err := createNetwork(cli, app.Name)
	if err != nil {
		return err
	}

	services, err := swarmSpec.GetServiceSpec(app.Name, networkID)
	if err != nil {
		return err
	}
	slog.Info("Get services from the source schema", "number of services found", len(services))

	// find the service if already exists
	allServicesRunning, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		return err
	}

	for _, service := range services {
		repo, found := repository.FindRepo(service.TaskTemplate.ContainerSpec.Image)
		auth := ""

		if !found {
			slog.Error("Repository not found in private image registries")
		} else {
			authString, err := repo.GetRegistryAuth()
			if err != nil {
				slog.Error(err.Error())
			} else {
				auth = authString
			}
		}

		// Checking if docker image is pullabel, if not then making the app health degraded.
		go func(cli *client.Client, a *Application) {
			// docker will not work if image is not reacheble\
			_, err = cli.ImagePull(context.TODO(), service.TaskTemplate.ContainerSpec.Image, types.ImagePullOptions{
				RegistryAuth: auth,
			})

			if err != nil {
				slog.Error("Failed to pull docker image, registry auth is required")
				a.Health = Degraded
			}
		}(cli, app)

		// check if already exists then only update
		if svc, exists := checkServiceAlreadyExist(service.Name, &allServicesRunning); exists {
			slog.Info("Service already running", "name", service.Name)
			res, err := cli.ServiceUpdate(context.Background(), svc.ID, svc.Version, service, types.ServiceUpdateOptions{
				EncodedRegistryAuth: auth,
			})
			if err != nil {
				app.Health = Degraded
				slog.Error("Not able to update a running service", "error", err.Error())
				return err
			}
			if len(res.Warnings) != 0 {
				slog.Warn("New Service update give warnings", "warnings", res.Warnings)
			}

			app.LastSyncedAt = time.Now()
			continue
		}

		slog.Info("Creating new service")
		res, err := cli.ServiceCreate(context.Background(), service, types.ServiceCreateOptions{
			EncodedRegistryAuth: auth,
		})
		if err != nil {
			app.Health = Degraded
			slog.Error("Not able to create a new service", "error", err.Error())
			return err
		}

		if len(res.Warnings) != 0 {
			slog.Warn("New Service Create give warnings", "warnings", res.Warnings)
		}

		app.LastSyncedAt = time.Now()
	}

	app.LiveState = targetState
	return nil
}

// SyncStatus Check if LiveState = TargetState
//
// Whether or not the live state matches the target state.
// Is the deployed application the same as Git says it should be?
func (app *Application) SyncStatus(targetState string) bool {
	return app.LiveState == targetState
}

func checkServiceAlreadyExist(serviceName string, allServices *[]swarm.Service) (swarm.Service, bool) {
	for _, svc := range *allServices {
		if svc.Spec.Name == serviceName {
			return svc, true
		}
	}
	return swarm.Service{}, false
}

func createNetwork(cli *client.Client, appName string) (string, error) {
	slog.Info("Creating network")
	networkName := appName + "_default"

	nets, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return "", err
	}

	for _, network := range nets {
		if network.Name == networkName {
			slog.Info("Network already exists")
			return network.ID, nil
		}
	}

	net, err := cli.NetworkCreate(context.Background(), networkName, types.NetworkCreate{
		Scope: "swarm",
		Labels: map[string]string{
			"com.docker.stack.namespace": appName,
		},
		Driver: "overlay",
	})

	slog.Info("Created network", "id", net.ID)

	if err != nil {
		return "", err
	}

	if net.Warning != "" {
		slog.Warn(net.Warning)
	}

	return net.ID, nil
}
