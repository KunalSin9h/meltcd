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

package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/meltred/meltcd/internal/core/application"
	"github.com/meltred/meltcd/server"
	api "github.com/meltred/meltcd/server/api/app"
	"github.com/meltred/meltcd/util"
	"github.com/spf13/cobra"
)

func CreateNewApplication(cmd *cobra.Command, args []string) error {
	spec, err := getSpecFromData(cmd, args)
	if err != nil {
		return err
	}

	app := application.New(spec)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(app); err != nil {
		return err
	}

	req, client, err := server.HTTPRequestWithBearerToken(http.MethodPost, fmt.Sprintf("%s/api/apps", util.GetServer()), buf, true)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return server.ReadAuthError(res.Body)
	}

	var resPayload api.GlobalResponse
	if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(resPayload.Message)
	}

	util.Info(resPayload.Message)
	return nil
}

func getSpecFromData(cmd *cobra.Command, args []string) (application.Spec, error) {
	var spec application.Spec

	if len(args) == 0 {
		util.Info("Application with Specification file")
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
		// menu using arguments
		util.Info("Creating application using arguments")
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
