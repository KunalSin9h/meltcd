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
	"meltred/meltcd/internal/core/application"
	"net/http"

	"github.com/spf13/cobra"
)

func createNewApplication(cmd *cobra.Command, args []string) error {
	var spec application.Spec

	if len(args) == 0 {
		info("Creating application with Specification file")
		// Creating application without application name
		// means using a file
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}
		spec, err = application.ParseSpecFromFile(file)
		if err != nil {
			return err
		}
	} else {
		// creating application with application name
		// means using arguments
		info("Creating application using arguments")
		name := args[0]

		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			return err
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}

		refresh, _ := cmd.Flags().GetString("refresh")
		revision, _ := cmd.Flags().GetString("revision")

		spec, err = application.ParseSpecFromValue(name, repo, revision, path, refresh)
		if err != nil {
			return err
		}
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
	// TODO Extract fiber.Error from the response
	if res.StatusCode != 202 {
		errorMsg("server not respond with 202")
		return errors.New("something went wrong")
	}

	info("New Application created")
	return nil
}
