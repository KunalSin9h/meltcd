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

package spec

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/swarm"
)

type DockerSwarm struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Networks map[string]Network `yaml:"networks"`
	Volumes  map[string]Volume  `yaml:"volumes"`
}

type Service struct {
	Build       string            `yaml:"build"`
	Image       string            `yaml:"image"`
	Ports       []string          `yaml:"ports"`
	Deploy      Deploy            `yaml:"deploy"`
	Environment map[string]string `yaml:"environment"`
	EnvFile     []string          `yaml:"env_file"`
	Volumes     []string          `yaml:"volumes"`
	Networks    []string          `yaml:"networks"`
}

type Deploy struct {
	Mode     string `yaml:"mode"`
	Replicas uint64 `yaml:"replicas"`
}

type Network struct {
	Name   string            `yaml:"name"`
	Driver string            `yaml:"driver"`
	Ipam   map[string]string `yaml:"ipam"`
}

type Volume struct {
	Name       string            `yaml:"name"`
	Driver     string            `yaml:"driver"`
	DriverOpts map[string]string `yaml:"driver_opts"`
	Labels     []string          `yaml:"labels"`
	Options    map[string]string `yaml:"options"`
}

func (d *DockerSwarm) GetServiceSpec(appName string) ([]swarm.ServiceSpec, error) {
	log.Info("Getting service spec for app", "app name", appName)

	specs := make([]swarm.ServiceSpec, 0)

	for serviceName, spec := range d.Services {
		spec := spec
		log.Info("Making serviceSpec for service", "service_name", serviceName)

		var targetSpec swarm.ServiceSpec
		targetSpec.Name = appName + "_" + serviceName
		targetSpec.TaskTemplate = swarm.TaskSpec{
			ContainerSpec: &swarm.ContainerSpec{
				Image: spec.Image,
			},
		}

		for _, envFile := range spec.EnvFile {
			envVars, err := getEnvVars(envFile)
			if err != nil {
				return []swarm.ServiceSpec{}, err
			}

			for k, v := range envVars {
				targetSpec.TaskTemplate.ContainerSpec.Env = append(targetSpec.TaskTemplate.ContainerSpec.Env, k+"="+v)
			}
		}

		for k, v := range spec.Environment {
			targetSpec.TaskTemplate.ContainerSpec.Env = append(targetSpec.TaskTemplate.ContainerSpec.Env, k+"="+v)
		}

		for _, m := range spec.Volumes {
			tokens := strings.Split(m, ":")
			if len(tokens) != 2 {
				log.Error("Volumes are not split on : in 2", "tokens", tokens)
				os.Exit(1)
			}

			targetSpec.TaskTemplate.ContainerSpec.Mounts = append(targetSpec.TaskTemplate.ContainerSpec.Mounts, mount.Mount{
				Type:   mount.TypeVolume,
				Source: tokens[0],
				Target: tokens[1],
			})
		}

		if spec.Deploy.Mode == "replicated" {
			targetSpec.Mode.Replicated = &swarm.ReplicatedService{
				Replicas: &spec.Deploy.Replicas,
			}
		} else if spec.Deploy.Mode == "global" {
			targetSpec.Mode.Global = &swarm.GlobalService{}
		}

		var ports []swarm.PortConfig
		for _, port := range spec.Ports {
			tokens := strings.Split(port, ":")
			if len(tokens) != 2 {
				log.Error("ports are not split on : in 2", "tokens", tokens)
				os.Exit(1)
			}
			target, _ := strconv.Atoi(tokens[1])
			publish, _ := strconv.Atoi(tokens[0])

			ports = append(ports, swarm.PortConfig{
				Protocol:      "tcp",
				TargetPort:    uint32(target),
				PublishedPort: uint32(publish),
				PublishMode:   swarm.PortConfigPublishModeIngress,
			})
		}

		targetSpec.EndpointSpec = &swarm.EndpointSpec{
			Ports: ports,
		}

		log.Info("Adding serviceSpec for service in allServiceArray", "service_name", serviceName)
		specs = append(specs, targetSpec)
	}

	return specs, nil
}

func getEnvVars(fileName string) (map[string]string, error) {
	result := make(map[string]string)

	fileName, err := normalizeFilePath(fileName)
	if err != nil {
		return map[string]string{}, err
	}

	fileData, err := os.Open(fileName)
	if err != nil {
		return map[string]string{}, err
	}
	defer fileData.Close()

	scanner := bufio.NewScanner(fileData)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.SplitN(line, "=", 2)

		if len(tokens) == 2 {
			key := strings.TrimSpace(tokens[0])
			value := strings.TrimSpace(tokens[1])
			value = strings.ReplaceAll(value, "\"", "")

			result[key] = value
		}
	}

	return result, err
}

func normalizeFilePath(fileName string) (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	username := currentUser.Username

	fileName = strings.ReplaceAll(fileName, "~", fmt.Sprintf("/home/%s", username))

	absFilePath, err := filepath.Abs(fileName)
	if err != nil {
		return "", err
	}

	return absFilePath, nil
}
