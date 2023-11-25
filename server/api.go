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

package server

import (
	"errors"
	"meltred/meltcd/internal/core"
	"meltred/meltcd/internal/core/application"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var app application.Application

	if err := c.BodyParser(&app); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "body parsing error")
	}

	if err := core.Register(&app); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendStatus(http.StatusAccepted)
}

func Update(c *fiber.Ctx) error {
	var app application.Application

	if err := c.BodyParser(&app); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "body parsing error")
	}

	if err := core.Update(&app); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.SendStatus(http.StatusAccepted)
}

func Details(c *fiber.Ctx) error {
	appName := c.Params("app_name")
	if appName == "" {
		return errors.New("application name (app_name) missing in querystring")
	}

	details, err := core.Details(appName)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(details)
}

type AppList struct {
	Data []AppStatus `json:"data"`
}

type AppStatus struct {
	Name   string `json:"name"`
	Health string `json:"health"`
}

func AllApplications(c *fiber.Ctx) error {
	status := core.List()

	var res AppList

	for k, v := range status {
		res.Data = append(res.Data, AppStatus{
			Name:   k,
			Health: v,
		})
	}

	return c.Status(200).JSON(res)
}

func Refresh(c *fiber.Ctx) error {
	appName := c.Params("app_name")

	if err := core.Refresh(appName); err != nil {
		return err
	}

	return c.SendStatus(200)
}
