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

package core

import (
	"fmt"
	"os"
	"path"

	"log/slog"

	"github.com/kunalsin9h/meltcd/internal/core/auth"
	"github.com/kunalsin9h/meltcd/internal/core/repository"
)

const MELTCD_DIR = ".meltcd"                         //nolint
const MELTCD_APPLICATIONS_FILE = "applications.json" //nolint
const MELTCD_REPOSITORY_FILE = "repositories.json"   //nolint
const MELTCD_AUTH_FILE = "auth.json"                 //nolint
const MELTCD_ACCESS_TOKEN = "access_token.txt"       //nolint
const MELTCD_LOG_FILE = "general.log"                //nolint

// Setup will setup require
// settings to make use of MeltCD
// like setting up admin password in docker secret
// setting up docker volume for persistent storage
//
// fill  the Applications from the stored file
//
// initialize a new docker client
func Setup() error {
	return meltcdState()
}

func meltcdState() error {
	applicationsFile := getAppFile()
	repositoryFile := getRepositoryFile()
	authFile := getAuthFile()
	accessTokenFile := getAccessTokenFile()

	// When creating a fresh auth file (db) insert admin:admin username and password
	_, err := os.Stat(authFile)
	if err != nil {
		slog.Info(fmt.Sprintf("Creating file: %s\n", authFile))
		_, err = os.Create(authFile)
		if err != nil {
			return err
		}
		slog.Info("Creating default user", "username", "admin", "password", "admin")
		if err := auth.InsertUser("admin", "admin", auth.Admin); err != nil {
			slog.Error("Failed to create default user")
			return err
		}
	}

	for _, f := range []string{applicationsFile, repositoryFile, accessTokenFile} {
		_, err = os.Stat(f)
		if err != nil {
			slog.Info(fmt.Sprintf("Creating file: %s\n", f))
			_, err = os.Create(f)
			if err != nil {
				return err
			}
		}
	}

	appData, err := os.ReadFile(applicationsFile)
	if err != nil {
		return err
	}

	if err := loadRegistryData(&appData); err != nil {
		slog.Warn("Application state file is empty", "error", err.Error())
	}

	repoData, err := os.ReadFile(repositoryFile)
	if err != nil {
		return err
	}

	if err := repository.LoadData(&repoData); err != nil {
		slog.Warn("Repository state file is empty", "error", err.Error())
	}

	authData, err := os.ReadFile(authFile)
	if err != nil {
		return err
	}

	if err := auth.LoadUsers(&authData); err != nil {
		slog.Warn("Auth file is empty", "error", err.Error())
	}

	return nil
}

func ShutDown() error {
	appFile := getAppFile()

	appData, err := getRegistryData()
	if err != nil {
		return err
	}

	if err := os.WriteFile(appFile, appData, os.ModePerm); err != nil {
		return err
	}

	repoFile := getRepositoryFile()

	repoData, err := repository.GetData()
	if err != nil {
		return err
	}

	if err := os.WriteFile(repoFile, repoData, os.ModePerm); err != nil {
		return err
	}

	authFile := getAuthFile()

	authData, err := auth.GetUsers()
	if err != nil {
		return err
	}

	return os.WriteFile(authFile, *authData, os.ModePerm)
}

func getMeltcdDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Warn("failed to get home dir (using default \".\")")
		return "."
	}

	meltcdDir := path.Join(home, MELTCD_DIR)

	return meltcdDir
}

func StoreAccessToken(token string) error {
	file := getAccessTokenFile()
	f, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = f.WriteString(token)
	if err != nil {
		return err
	}

	return nil
}

func GetAccessToken() string {
	file := getAccessTokenFile()
	data, err := os.ReadFile(file)
	if err != nil {
		return ""
	}

	return string(data)
}

func getAppFile() string {
	meltcdDir := getMeltcdDir()
	return path.Join(meltcdDir, MELTCD_APPLICATIONS_FILE)
}

func getRepositoryFile() string {
	meltcdDir := getMeltcdDir()
	return path.Join(meltcdDir, MELTCD_REPOSITORY_FILE)
}

func getAuthFile() string {
	meltcdDir := getMeltcdDir()
	return path.Join(meltcdDir, MELTCD_AUTH_FILE)
}

func getAccessTokenFile() string {
	meltcdDir := getMeltcdDir()
	return path.Join(meltcdDir, MELTCD_ACCESS_TOKEN)
}

func getLogFile() string {
	meltcdDir := getMeltcdDir()
	return path.Join(meltcdDir, MELTCD_LOG_FILE)
}
