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

package api

import (
	"net/http"

	"github.com/meltred/meltcd/internal/core"
	"github.com/meltred/meltcd/internal/core/application"

	"github.com/gofiber/fiber/v2"
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

	if err := core.Register(&app); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(GlobalResponse{
		Message: "Application registered successfully",
	})
}

// Update godoc
//
//	@summary	Update an application
//	@tags		Apps
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

// Details godoc
//
//	@summary	Get details of an application
//	@tags		Apps
//	@param		app_name	path	string	true	"Application name"
//	@produce	json
//	@success	200	{object}	application.Application
//	@failure	400	{object}	GlobalResponse
//	@failure	500	{object}	GlobalResponse
//	@router		/apps/{app_name} [get]
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

// AllApplications godoc
//
//	@summary	Get a list all applications created
//	@tags		Apps
//	@success	200	{object}	core.AppList
//	@router		/apps [get]
func AllApplications(c *fiber.Ctx) error {
	status := core.List()
	return c.Status(200).JSON(status)
}

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

// Remove godoc
//
//	@summary	Remove an application
//	@tags		Apps
//	@param		app_name	path	string	true	"Application name"
//	@success	200
//	@failure	500	{object}	GlobalResponse
//	@router		/apps/{app_name} [delete]
func Remove(c *fiber.Ctx) error {
	appName := c.Params("app_name")

	if err := core.RemoveApplication(appName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.SendStatus(200)
}
