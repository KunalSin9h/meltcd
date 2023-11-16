/*
Copyright 2023 - PRESENT Meltred Pvt. Ltd.

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
	"log"
	"meltred/meltcd/version"

	"github.com/spf13/cobra"
)

// NewApplication creates a new cli app
// This cli app can be used to start the api server
// as well as a client
func NewApplication() *cobra.Command {
	// Log time, date, and file name
	log.SetFlags(log.Ldate | log.Lshortfile)

	rootCmd := &cobra.Command{
		Use:           "metlcd",
		Short:         "ArgoCD like Continuous Deployment for Docker Swarm",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Version: version.Version,
	}

	cobra.EnableCommandSorting = false

	// Server / Start API
	serveCmd := &cobra.Command{
		Use:     "serve",
		Aliases: []string{"start"},
		Short:   "Start MeltCD server",
		Args:    cobra.ExactArgs(0),
		RunE:    RunServer,
	}

	serveCmd.Flags().Bool("verbose", false, "verbose is used to get extra logs/info about process")

	rootCmd.AddCommand(serveCmd)

	return rootCmd
}
