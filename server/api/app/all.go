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

package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core"
)

type GlobalResponse struct {
	Message string `json:"message"`
}

// AllApplications godoc
//
//	@summary	Get a list all applications created
//	@tags		Apps
//	@Security	ApiKeyAuth || cookies
//	@success	200	{object}	core.AppList
//	@router		/apps [get]
func AllApplications(c *fiber.Ctx) error {
	status := core.List()
	return c.Status(200).JSON(status)
}
