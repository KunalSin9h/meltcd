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

package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kunalsin9h/meltcd/internal/core/application"
	"github.com/kunalsin9h/meltcd/server"
	api "github.com/kunalsin9h/meltcd/server/api/app"
	"github.com/kunalsin9h/meltcd/util"
	"github.com/spf13/cobra"
)

func UpdateExistingApplication(cmd *cobra.Command, args []string) error {
	spec, err := getSpecFromData(cmd, args)
	if err != nil {
		return err
	}

	app := application.New(spec)

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(app); err != nil {
		return err
	}

	req, client, err := server.HTTPRequestWithBearerToken(http.MethodPut, fmt.Sprintf("%s/api/apps", util.GetServer()), buf, true)

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

	if res.StatusCode != http.StatusAccepted {
		var resPayload api.GlobalResponse
		if err := json.NewDecoder(req.Body).Decode(&resPayload); err != nil {
			return err
		}
		return errors.New(resPayload.Message)
	}

	util.Info("Application updated")
	return nil
}
