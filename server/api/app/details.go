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

package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kunalsin9h/meltcd/internal/core"
)

// Details godoc
// @summary	Get details of an application
// @tags		Apps
// @Security	ApiKeyAuth || cookies
// @param		app_name	path	string	true	"Application name"
// @produce	json
// @success	200	{object}	application.Application
// @failure	400	{object}	GlobalResponse
// @failure	500	{object}	GlobalResponse
// @router		/apps/{app_name} [get]
func Details(c *fiber.Ctx) error {
	appName := c.Params("app_name")

	if appName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: "application name (app_name) missing in querystring",
		})
	}

	details, err := core.Details(appName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(details)
}
