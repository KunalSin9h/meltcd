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
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core"
	"github.com/meltred/meltcd/internal/core/application"
)

// Update godoc
//
//	@summary	Update an application
//	@tags		Apps
//	@Security	ApiKeyAuth || cookies
//	@accept		json
//	@produce	json
//	@param		request	body	application.Application	true	"Application body"
//	@success	202
//	@failure	400	{object}	GlobalResponse
//	@failure	500	{object}	GlobalResponse
//	@router		/apps [put]
func Update(c *fiber.Ctx) error {
	var app application.Application

	if err := c.BodyParser(&app); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: "Failed to parse request body",
		})
	}

	if err := core.Update(&app); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.SendStatus(http.StatusAccepted)
}
