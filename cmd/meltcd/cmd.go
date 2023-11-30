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
	"log"

	"github.com/meltred/meltcd/version"

	"github.com/spf13/cobra"
)

// NewApplication creates a new cli app
// This cli app can be used to start the api server
// as well as a client
func NewCLI() *cobra.Command {
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

	// Application
	appCmd := &cobra.Command{
		Use:   "app",
		Short: "Work with Applications",
	}

	appCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Args:  cobra.RangeArgs(0, 1),
		RunE:  createNewApplication,
	}

	appCreateCmd.Flags().String("repo", "", "The git repository where the service file is hosted")
	appCreateCmd.Flags().String("revision", "HEAD", "The git repository revision")
	appCreateCmd.Flags().String("path", "", "The path to service file")
	appCreateCmd.Flags().String("refresh", "3m0s", "The refresh time for sync")
	appCreateCmd.Flags().String("file", "", "Application schema file")

	appUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update existing application",
		Args:  cobra.RangeArgs(0, 1),
		RunE:  updateExistingApplication,
	}

	appUpdateCmd.Flags().String("repo", "", "The git repository where the service file is hosted")
	appUpdateCmd.Flags().String("revision", "HEAD", "The git repository revision")
	appUpdateCmd.Flags().String("path", "", "The path to service file")
	appUpdateCmd.Flags().String("refresh", "3m0s", "The refresh time for sync")
	appUpdateCmd.Flags().String("file", "", "Application schema file")

	appGetCmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"inspect"},
		Short:   "Get details about the application",
		Args:    cobra.ExactArgs(1),
		RunE:    getDetailsAboutApplication,
	}

	appListCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Get all the applications registered",
		Args:    cobra.ExactArgs(0),
		RunE:    getAllApplications,
	}

	appRefreshCmd := &cobra.Command{
		Use:     "refresh",
		Aliases: []string{"sync"},
		Short:   "Force refresh (synchronize) application",
		Args:    cobra.ExactArgs(1),
		RunE:    refreshApplication,
	}

	appRemoveCmd := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"rm"},
		Short:   "Remove Application",
		Args:    cobra.ExactArgs(1),
		RunE:    removeApplication,
	}

	appCmd.AddCommand(appCreateCmd)
	appCmd.AddCommand(appUpdateCmd)
	appCmd.AddCommand(appGetCmd)
	appCmd.AddCommand(appListCmd)
	appCmd.AddCommand(appRefreshCmd)
	appCmd.AddCommand(appRemoveCmd)

	// meltcd repo
	repoCmd := &cobra.Command{
		Use:     "repo",
		Aliases: []string{"repository"},
		Short:   "Working with private git repository",
	}

	// meltcd  repo add https://github.com/... --username "" --password ""
	repoAddCmd := &cobra.Command{
		Use:   "add REPO_URL",
		Short: "Add a private git repository",
		Args:  cobra.ExactArgs(1), // the git repo url
		RunE:  addPrivateGitRepository,
	}

	repoAddCmd.Flags().String("username", "", "username for basic auth")
	repoAddCmd.MarkFlagRequired("username")
	repoAddCmd.Flags().String("password", "", "password for basic auth")
	repoAddCmd.MarkFlagRequired("password")

	repoCmd.AddCommand(repoAddCmd)

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(appCmd)
	rootCmd.AddCommand(repoCmd)

	return rootCmd
}
