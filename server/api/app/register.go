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
	"github.com/meltred/meltcd/internal/core/application"
)

// Register godoc
//
//	@summary	Create a new application
//	@tags		Apps
//	@accept		json
//	@produce	json
//	@param		request	body		application.Application	true	"Application body"
//	@success	200		{object}	GlobalResponse
//	@failure	400		{object}	GlobalResponse
//	@failure	500		{object}	GlobalResponse
//	@router		/apps [post]
func Register(c *fiber.Ctx) error {
	var app application.Application

	if err := c.BodyParser(&app); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: "Failed to pase request body",
		})
	}

	// clearing the current state, so it can be fetch again
	app.LiveState = ""

	if err := core.Register(&app); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(GlobalResponse{
		Message: "Application registered successfully",
	})
}
