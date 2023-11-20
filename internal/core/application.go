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

import "github.com/docker/docker/api/types/swarm"

type Application struct {
	Name         string            `json:"name"`
	RefreshTimer uint64            `json:"sync_timer"` // Timer to check for Sync
	Health       ApplicationHealth `json:"health"`
	LiveState    swarm.ServiceSpec `json:"live_state"`
}

type ApplicationHealth int

const (
	Healthy ApplicationHealth = iota
	Progressing
	Degraded
	Suspended
)

// SyncStatus Check if LiveState = TargetState
//
// Whether or not the live state matches the target state.
// Is the deployed application the same as Git says it should be?
func (app *Application) SyncStatus(targetState swarm.ServiceSpec) bool {
	//TODO
	return false
}

// Sync
// The process of making an application move to its target state.
// E.g. by applying changes to a docker swarm cluster.
func (app *Application) Sync(targetState swarm.ServiceSpec) error {
	//TODO
	return nil
}
