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

package meltcd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kunalsin9h/meltcd/server"
	"github.com/kunalsin9h/meltcd/server/api/app"
	"github.com/kunalsin9h/meltcd/server/api/repo"
	"github.com/kunalsin9h/meltcd/util"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func addPrivateRepository(cmd *cobra.Command, args []string) error {
	repoName := args[0]

	git, _ := cmd.Flags().GetBool("git")
	image, _ := cmd.Flags().GetBool("image")
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")

	payload := repo.PrivateRepoDetails{
		Username: username,
		Password: password,
	}

	// if not git then if image is also false, then default is git
	if git || (!git && !image) {
		repoName, _ = strings.CutSuffix(repoName, "/")
		payload.URL = repoName
	} else if image {
		payload.ImageRef = repoName
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}

	req, client, err := server.HTTPRequestWithBearerToken(http.MethodPost, fmt.Sprintf("%s/api/repo", util.GetServer()), buf, true)
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

	var resBody app.GlobalResponse
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return err
	}

	if res.StatusCode != fiber.StatusAccepted {
		return errors.New(resBody.Message)
	}

	fmt.Println(resBody.Message)
	return nil
}

func getAllRepoAdded(_ *cobra.Command, _ []string) error {
	req, client, err := server.HTTPRequestWithBearerToken(http.MethodGet, fmt.Sprintf("%s/api/repo", util.GetServer()), nil, false)
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

	var resData repo.ListData
	if err := json.NewDecoder(res.Body).Decode(&resData); err != nil {
		return err
	}

	tableRepositoryGit := table.New("S.NO", "Repository URL", "Reachable")
	tableRepositoryGit.WithHeaderFormatter(util.HeaderFmt).WithFirstColumnFormatter(util.ColumnFmt)

	tableContainerImage := table.New("S.NO", "Container Image", "Reachable")
	tableContainerImage.WithHeaderFormatter(util.HeaderFmt).WithFirstColumnFormatter(util.ColumnFmt)

	i, j := 1, 1
	for _, repo := range resData.Data {
		if repo.URL != "" {
			tableRepositoryGit.AddRow(i, repo.URL, repo.Reachable)
			i++
		} else {
			tableContainerImage.AddRow(j, repo.ImageRef, repo.Reachable)
			j++
		}
	}

	tableRepositoryGit.Print()
	fmt.Println()
	tableContainerImage.Print()
	return nil
}

func removePrivateRepo(_ *cobra.Command, args []string) error {
	repoURL := args[0]
	repoURL, _ = strings.CutSuffix(repoURL, "/")

	payload := repo.RemovePayload{
		Repo: repoURL,
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}

	request, client, err := server.HTTPRequestWithBearerToken(http.MethodDelete, fmt.Sprintf("%s/api/repo", util.GetServer()), buf, true)
	if err != nil {
		return err
	}

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return server.ReadAuthError(res.Body)
	}

	var data app.GlobalResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	fmt.Println(data.Message)
	return nil
}

func updatePrivateRepo(cmd *cobra.Command, args []string) error {
	repoName := args[0]

	git, _ := cmd.Flags().GetBool("git")
	image, _ := cmd.Flags().GetBool("image")
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")

	gitRepo, imageRepo := "", ""

	if git {
		gitRepo = repoName
	}

	if image {
		imageRepo = repoName
	}

	payload := repo.PrivateRepoDetails{
		URL:      gitRepo,
		ImageRef: imageRepo,
		Username: username,
		Password: password,
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}

	req, client, err := server.HTTPRequestWithBearerToken(http.MethodPut, fmt.Sprintf("%s/api/repo", util.GetServer()), buf, true)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusUnauthorized {
		return server.ReadAuthError(res.Body)
	}

	defer res.Body.Close()

	var data app.GlobalResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return err
	}

	fmt.Println(data.Message)
	return nil
}
