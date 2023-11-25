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
	"net"
	"os"
	"strings"

	"github.com/meltred/meltcd/server"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func RunServer(cmd *cobra.Command, _ []string) error {
	baseURL := os.Getenv("MELTCD_HOST")
	verbose, _ := cmd.Flags().GetBool("verbose")

	host, port, err := net.SplitHostPort(baseURL)
	if err != nil {
		log.Warn(err)

		host, port = "127.0.0.1", "11771"
		if ip := net.ParseIP(strings.Trim(baseURL, "[]")); ip != nil {
			host = ip.String()
		}
	}

	ln, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}

	origins := os.Getenv("MELTCD_ORIGINS")

	return server.Serve(ln, origins, verbose)
}
