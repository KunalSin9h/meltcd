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

// Refresh godoc
//
//	@summary	Refresh/Synchronize an application
//	@tags		Apps
//	@param		app_name	path	string	true	"Application name"
//	@success	200
//	@failure	500	{object}	GlobalResponse
//	@router		/apps/{app_name}/refresh [post]
func Refresh(c *fiber.Ctx) error {
	appName := c.Params("app_name")

	if err := core.Refresh(appName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.SendStatus(200)
}
