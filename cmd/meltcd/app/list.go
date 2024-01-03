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
	"fmt"
	"net/http"

	"github.com/meltred/meltcd/internal/core"
	"github.com/meltred/meltcd/server"
	"github.com/meltred/meltcd/util"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func GetAllApplications(_ *cobra.Command, _ []string) error {
	req, client, err := server.HTTPRequestWithBearerToken(http.MethodGet, fmt.Sprintf("%s/api/apps", util.GetServer()), nil, false)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return server.ReadAuthError(res.Body)
	}

	var resPayload core.AppList
	if err := json.NewDecoder(res.Body).Decode(&resPayload); err != nil {
		return err
	}

	table := table.New("S.NO", "Name", "Health", "Last Synced At", "Created At", "Updated At")
	table.WithHeaderFormatter(util.HeaderFmt).WithFirstColumnFormatter(util.ColumnFmt)

	for _, v := range resPayload.Data {
		table.AddRow(v.ID, v.Name, v.Health, util.GetSinceTime(v.LastSyncedAt), util.GetSinceTime(v.UpdatedAT), util.GetSinceTime(v.CreatedAt))
	}

	table.Print()
	return nil
}
