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
	"log"

	"github.com/kunalsin9h/meltcd/cmd/meltcd/app"
	"github.com/kunalsin9h/meltcd/version"

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

	rootCmd.AddCommand(serveCmd)

	// auth commands
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login user",
		Args:  cobra.ExactArgs(0),
		RunE:  LoginUser,
	}

	loginCmd.Flags().Bool("show-token", false, "Get the Access token when logged in successfully")

	rootCmd.AddCommand(loginCmd)

	// Application
	appCmd := &cobra.Command{
		Use:   "app",
		Short: "Work with Applications",
	}

	appCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Args:  cobra.RangeArgs(0, 1),
		RunE:  app.CreateNewApplication,
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
		RunE:  app.UpdateExistingApplication,
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
		RunE:    app.GetDetailsAboutApplication,
	}

	appListCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Get all the applications registered",
		Args:    cobra.ExactArgs(0),
		RunE:    app.GetAllApplications,
	}

	appRefreshCmd := &cobra.Command{
		Use:     "refresh",
		Aliases: []string{"sync"},
		Short:   "Force refresh (synchronize) application",
		Args:    cobra.ExactArgs(1),
		RunE:    app.RefreshApplication,
	}

	appRemoveCmd := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"rm"},
		Short:   "Remove Application",
		Args:    cobra.ExactArgs(1),
		RunE:    app.RemoveApplication,
	}

	appRecreateCmd := &cobra.Command{
		Use:     "recreate [APP_NAME]",
		Aliases: []string{"rc"},
		Short:   "Recreate Application",
		Args:    cobra.ExactArgs(1),
		RunE:    app.RecreateApplication,
	}

	appCmd.AddCommand(appCreateCmd)
	appCmd.AddCommand(appUpdateCmd)
	appCmd.AddCommand(appGetCmd)
	appCmd.AddCommand(appListCmd)
	appCmd.AddCommand(appRefreshCmd)
	appCmd.AddCommand(appRemoveCmd)
	appCmd.AddCommand(appRecreateCmd)

	rootCmd.AddCommand(appCmd)

	// meltcd repo
	repoCmd := &cobra.Command{
		Use:     "repo",
		Aliases: []string{"repository"},
		Short:   "Working with private git repository",
	}

	// meltcd  repo add https://github.com/... --username "" --password ""
	repoAddCmd := &cobra.Command{
		Use:   "add REPO",
		Short: "Add a private git repository (--git) or image registry (--image)",
		Args:  cobra.ExactArgs(1), // the git repo url
		RunE:  addPrivateRepository,
	}

	repoAddCmd.Flags().Bool("git", false, "if private repo is a git repository")
	repoAddCmd.Flags().Bool("image", false, "if private repo is a docker image")
	repoAddCmd.Flags().String("username", "", "username for basic auth")
	repoAddCmd.MarkFlagRequired("username")
	repoAddCmd.Flags().String("password", "", "password for basic auth")
	repoAddCmd.MarkFlagRequired("password")

	repoListCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all added repositories",
		RunE:    getAllRepoAdded,
	}

	repoRemoveCmd := &cobra.Command{
		Use:     "remove REPO_URL",
		Aliases: []string{"rm"},
		Short:   "Remove private repository",
		Args:    cobra.ExactArgs(1), // repo-url
		RunE:    removePrivateRepo,
	}

	repoUpdateCmd := &cobra.Command{
		Use:   "update REPO_URL",
		Short: "Update auth credentials for private repository",
		Args:  cobra.ExactArgs(1), // the git repo url
		RunE:  updatePrivateRepo,
	}

	repoUpdateCmd.Flags().Bool("git", false, "if private repo is a git repository")
	repoUpdateCmd.Flags().Bool("image", false, "if private repo is a docker image")
	repoUpdateCmd.Flags().String("username", "", "username for basic auth")
	repoUpdateCmd.MarkFlagRequired("username")
	repoUpdateCmd.Flags().String("password", "", "password for basic auth")
	repoUpdateCmd.MarkFlagRequired("password")

	repoCmd.AddCommand(repoAddCmd)
	repoCmd.AddCommand(repoListCmd)
	repoCmd.AddCommand(repoRemoveCmd)
	repoCmd.AddCommand(repoUpdateCmd)

	rootCmd.AddCommand(repoCmd)

	return rootCmd
}
