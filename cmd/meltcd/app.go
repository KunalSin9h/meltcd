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
	"net/http"

	"github.com/meltred/meltcd/internal/core"
	"github.com/meltred/meltcd/internal/core/application"
	"github.com/meltred/meltcd/server/api"
	"github.com/rodaine/table"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func getDetailsAboutApplication(_ *cobra.Command, args []string) error {
	appName := args[0]

	res, err := http.Get(fmt.Sprintf("%s/api/apps/%s", getServer(), appName))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		var resPayload api.GlobalResponse
		if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
			return err
		}
		return errors.New(resPayload.Message)
	}

	var resDada application.Application
	if err := json.NewDecoder(res.Body).Decode(&resDada); err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(resDada, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}

func createNewApplication(cmd *cobra.Command, args []string) error {
	spec, err := getSpecFromData(cmd, args)
	if err != nil {
		return err
	}

	app := application.New(spec)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(app); err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%s/api/apps", getServer()), "application/json", buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var resPayload api.GlobalResponse
	if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
		return err
	}
	if res.StatusCode != fiber.StatusOK {
		return errors.New(resPayload.Message)
	}

	info(resPayload.Message)
	return nil
}

func updateExistingApplication(cmd *cobra.Command, args []string) error {
	spec, err := getSpecFromData(cmd, args)
	if err != nil {
		return err
	}

	app := application.New(spec)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(app); err != nil {
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest(fiber.MethodPut, fmt.Sprintf("%s/api/apps", getServer()), buf)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != fiber.StatusAccepted {
		var resPayload api.GlobalResponse
		if err := json.NewDecoder(req.Body).Decode(&resPayload); err != nil {
			return err
		}
		return errors.New(resPayload.Message)
	}

	info("Application updated")
	return nil
}

func getAllApplications(_ *cobra.Command, _ []string) error {
	res, err := http.Get(fmt.Sprintf("%s/api/apps", getServer()))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != fiber.StatusOK {
		return errors.New("server does not respond with 200")
	}

	var resPayload core.AppList
	if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
		return err
	}

	table := table.New("S.NO", "Name", "Health", "Last Synced At", "Created At", "Updated At")
	table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, v := range resPayload.Data {
		table.AddRow(v.ID, v.Name, v.Health, getSinceTime(v.LastSyncedAt), getSinceTime(v.UpdatedAT), getSinceTime(v.CreatedAt))
	}

	table.Print()
	return nil
}

func getSpecFromData(cmd *cobra.Command, args []string) (application.Spec, error) {
	var spec application.Spec

	if len(args) == 0 {
		info("Application with Specification file")
		// Creating application without application name
		// means using a file

		// if user has specified --repo then he/she must have forgotten the app name
		// and come here
		repo, err := cmd.Flags().GetString("repo")
		if repo != "" && err == nil {
			return application.Spec{}, errors.New("missing application name")
		}

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

	res, err := http.Post(fmt.Sprintf("%s/api/apps/%s/refresh", getServer(), appName), "", nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != fiber.StatusOK {
		var resPayload api.GlobalResponse
		if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
			return err
		}
		return errors.New(resPayload.Message)
	}

	info("Application Synchronized")
	return nil
}

func removeApplication(_ *cobra.Command, args []string) error {
	appName := args[0]

	client := &http.Client{}
	req, err := http.NewRequest(fiber.MethodDelete, fmt.Sprintf("%s/api/apps/%s", getServer(), appName), nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != fiber.StatusOK {
		var resPayload api.GlobalResponse
		if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
			return err
		}
		return errors.New(resPayload.Message)
	}

	info("Application removed")
	return nil
}
