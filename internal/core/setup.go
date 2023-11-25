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

package core

import (
	"os"
	"path"

	"github.com/charmbracelet/log"
)

const MELTCD_DIR = ".meltcd"                         //nolint
const MELTCD_APPLICATIONS_FILE = "applications.json" //nolint

// Setup will setup require
// settings to make use of MeltCD
// like setting up admin password in docker secret
// setting up docker volume for persistent storage
//
// fill  the Applications from the volume
//
// initialize a new docker client
func Setup() error {
	return meltcdState()
}

func getMeltcdDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Warn("failed to get home dir (using default \".\")")
		return "."
	}

	meltcdDir := path.Join(home, MELTCD_DIR)

	return meltcdDir
}

func getAppFile() string {
	meltcdDir := getMeltcdDir()
	return path.Join(meltcdDir, MELTCD_APPLICATIONS_FILE)
}

func meltcdState() error {
	meltcdDir := getMeltcdDir()

	_, err := os.Stat(meltcdDir)
	if err != nil {
		log.Infof("Creating directory: %s", meltcdDir)

		err = os.Mkdir(meltcdDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	applicationsFile := path.Join(meltcdDir, MELTCD_APPLICATIONS_FILE)

	_, err = os.Stat(applicationsFile)
	if err != nil {
		log.Infof("Creating file: %s", applicationsFile)
		_, err = os.Create(applicationsFile)
		if err != nil {
			return err
		}
	}

	data, err := os.ReadFile(applicationsFile)
	if err != nil {
		return err
	}

	if err := loadRegistryData(&data); err != nil {
		log.Warn("Application state file is empty")
	}

	return nil
}

func ShutDown() error {
	appFile := getAppFile()

	data, err := getRegistryData()
	if err != nil {
		return err
	}

	return os.WriteFile(appFile, data, os.ModePerm)
}
