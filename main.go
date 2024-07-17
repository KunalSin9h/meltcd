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

package main

import (
	"context"

	"github.com/kunalsin9h/meltcd/cmd/meltcd"
	_ "github.com/kunalsin9h/meltcd/docs/swagger"

	"github.com/spf13/cobra"
)

// @title						Meltcd API
// @version					0.6
// @description				Argo-cd like GitDevOps Continuous Development platform for docker swarm.
// @host						localhost:1771
// @basePath					/api
// @schemes					http
// @license.name				Apache 2.0
// @license.url				https://github.com/kunalsin9h/meltcd/blob/main/LICENSE
// @externalDocs.description	Meltcd Docs
// @externalDocs.url			https://cd.kunalsin9h.com/docs
// @securityDefinitions.basic	BasicAuth
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						X-API-Key
func main() {
	cobra.CheckErr(meltcd.NewCLI().ExecuteContext(context.Background()))
}
