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

package meltcd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/meltred/meltcd/internal/core/application"
	"github.com/meltred/meltcd/server"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func getDetailsAboutApplication(_ *cobra.Command, args []string) error {
	appName := args[0]

	res, err := http.Get(fmt.Sprintf("%s/api/application/get/%s", getServer(), appName))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("server does not respond with 200")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func createNewApplication(cmd *cobra.Command, args []string) error {
	spec, err := getSpecFromData(cmd, args)
	if err != nil {
		return err
	}

	app := application.New(spec)
	payload, err := json.Marshal(app)
	if err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%s/api/application/register", getServer()), "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 202 {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(data))
	}

	info("New Application created")
	return nil
}

func updateExistingApplication(cmd *cobra.Command, args []string) error {
	spec, err := getSpecFromData(cmd, args)
	if err != nil {
		return err
	}

	app := application.New(spec)
	payload, err := json.Marshal(app)
	if err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%s/api/application/update", getServer()), "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 202 {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return errors.New(string(data))
	}

	info("Application updated")
	return nil
}

func getAllApplications(_ *cobra.Command, _ []string) error {
	res, err := http.Get(fmt.Sprintf("%s/api/application/get", getServer()))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("server does not respond with 200")
	}

	var resPayload server.AppList
	if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
		return err
	}

	for _, v := range resPayload.Data {
		fmt.Println(v.Name, v.Health)
	}

	return nil
}

func getSpecFromData(cmd *cobra.Command, args []string) (application.Spec, error) {
	var spec application.Spec

	if len(args) == 0 {
		info("Application with Specification file")
		// Creating application without application name
		// means using a file
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return application.Spec{}, err
		}
		spec, err = application.ParseSpecFromFile(file)
		if err != nil {
			return application.Spec{}, err
		}
	} else {
		// creating application with application name
		// means using arguments
		info("Creating application using arguments")
		name := args[0]

		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			return application.Spec{}, err
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return application.Spec{}, err
		}

		refresh, _ := cmd.Flags().GetString("refresh")
		revision, _ := cmd.Flags().GetString("revision")

		spec, err = application.ParseSpecFromValue(name, repo, revision, path, refresh)
		if err != nil {
			return application.Spec{}, err
		}
	}

	return spec, nil
}

func refreshApplication(_ *cobra.Command, args []string) error {
	appName := args[0]

	res, err := http.Post(fmt.Sprintf("%s/api/application/refresh/%s", getServer(), appName), "", nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != fiber.StatusOK {
		return errors.New("server does not respond with 200")
	}
	return nil
}
