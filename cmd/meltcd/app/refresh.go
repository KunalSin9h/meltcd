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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/meltred/meltcd/server"
	api "github.com/meltred/meltcd/server/api/app"
	"github.com/meltred/meltcd/util"
	"github.com/spf13/cobra"
)

func RefreshApplication(_ *cobra.Command, args []string) error {
	appName := args[0]

	req, client, err := server.HTTPRequestWithBearerToken(http.MethodPost, fmt.Sprintf("%s/api/apps/%s/refresh", util.GetServer(), appName), nil, false)
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

	if res.StatusCode != http.StatusOK {
		var resPayload api.GlobalResponse
		if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
			return err
		}
		return errors.New(resPayload.Message)
	}

	util.Info("Application Synchronized")
	return nil
}
