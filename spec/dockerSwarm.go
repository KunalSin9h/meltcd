package spec

import (
	"os"
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
	Name    string            `yaml:"name"`
	Driver  string            `yaml:"driver"`
	Options map[string]string `yaml:"options"`
}

func (d *DockerSwarm) GetServiceSpec(appName string) ([]swarm.ServiceSpec, error) {
	log.Info("Getting service spec for app", "app name", appName)

	var specs []swarm.ServiceSpec

	for serviceName, spec := range d.Services {
		log.Info("Making serviceSpec for service", "service_name", serviceName)

		var targetSpec swarm.ServiceSpec
		targetSpec.Name = appName + "_" + serviceName
		targetSpec.TaskTemplate = swarm.TaskSpec{
			ContainerSpec: &swarm.ContainerSpec{
				Image: spec.Image,
			},
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
				Type:   mount.TypeBind,
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
