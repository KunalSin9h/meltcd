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

	"github.com/meltred/meltcd/internal/core/application"
	api "github.com/meltred/meltcd/server/api/app"
	"github.com/meltred/meltcd/util"
	"github.com/spf13/cobra"
)

func GetDetailsAboutApplication(_ *cobra.Command, args []string) error {
	appName := args[0]

	res, err := http.Get(fmt.Sprintf("%s/api/apps/%s", util.GetServer(), appName))
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
