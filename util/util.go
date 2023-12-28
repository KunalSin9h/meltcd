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

package util

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var HeaderFmt = color.New(color.FgGreen, color.Underline).SprintfFunc()
var ColumnFmt = color.New(color.FgYellow).SprintfFunc()

func GetServer() string {
	server := "http://127.0.0.1:11771"
	serverEnvVar := os.Getenv("MELTCD_SERVER")
	if serverEnvVar != "" {
		server, _ = strings.CutSuffix(serverEnvVar, "/")
	}

	return server
}

func Info(text string, args ...any) { // nolint
	fmt.Printf(text+"\n", args...)
}

func errorMsg(text string, args ...any) { //nolint
	fmt.Printf(text+"\n", args...)
}

func GetSinceTime(t time.Time) string {
	elapsed := time.Since(t).Milliseconds()

	if elapsed == 0 {
		return "Just now"
	}

	sec := elapsed / 1000
	mins := sec / 60
	hrs := mins / 60
	days := hrs / 24
	weeks := days / 7
	months := weeks / 4
	year := months / 12

	if year > 0 {
		return fmt.Sprintf("%d year ago", year)
	}

	if months > 0 {
		return fmt.Sprintf("%d months ago", months)
	}

	if weeks > 0 {
		return fmt.Sprintf("%d weeks ago", weeks)
	}

	if days > 0 {
		return fmt.Sprintf("%d days ago", days)
	}

	if hrs > 0 {
		return fmt.Sprintf("%d hours ago", hrs)
	}

	if mins > 0 {
		return fmt.Sprintf("%d minutes ago", mins)
	}

	return "Just now"
}
