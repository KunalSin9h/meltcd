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
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/server/api"
	"github.com/spf13/cobra"
)

func addPrivateGitRepository(cmd *cobra.Command, args []string) error {
	repoURL := args[0]
	repoURL, _ = strings.CutSuffix(repoURL, "/")

	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")

	payload := api.PrivateRepoDetails{
		URL:      repoURL,
		Username: username,
		Password: password,
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%s/api/repo/add", getServer()), "application/json", buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var resBody api.GlobalResponse
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return err
	}

	if res.StatusCode != fiber.StatusAccepted {
		return errors.New(resBody.Message)
	}

	fmt.Println(resBody.Message)
	return nil
}
