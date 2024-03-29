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

package repo

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kunalsin9h/meltcd/internal/core/repository"
	"github.com/kunalsin9h/meltcd/server/api/app"
)

type PrivateRepoDetails struct {
	URL      string `json:"url"`
	ImageRef string `json:"image_ref"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Add godoc
//
//	@summary	Add a new repository
//	@tags		Repo
//	@Security	ApiKeyAuth || cookies
//	@accept		json
//	@produce	json
//	@param		request	body		PrivateRepoDetails	true	"Repository details"
//	@success	202		{object}	app.GlobalResponse
//	@failure	400		{object}	app.GlobalResponse
//	@failure	500		{object}	app.GlobalResponse
//	@router		/repo [post]
func Add(c *fiber.Ctx) error { // nolint:all
	var payload PrivateRepoDetails

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(app.GlobalResponse{
			Message: err.Error(),
		})
	}

	if (payload.URL == "" && payload.ImageRef == "") || payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(app.GlobalResponse{
			Message: "missing url or imageRef, username or password in request body",
		})
	}

	url, _ := strings.CutSuffix(payload.URL, "/")

	if err := repository.Add(url, payload.ImageRef, payload.Username, payload.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(app.GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(app.GlobalResponse{
		Message: "Added Repository",
	})
}

type ListData struct {
	Data []repository.RepoData `json:"data"`
}

// List godoc
//
//	@summary	Get a list all repositories
//	@Security	ApiKeyAuth || cookies
//	@tags		Repo
//	@produce	json
//	@success	200	{object}	ListData
//	@router		/repo [get]
func List(c *fiber.Ctx) error {
	list := repository.List()

	return c.Status(fiber.StatusOK).JSON(ListData{
		Data: list,
	})
}

type RemovePayload struct {
	Repo string `json:"repo"`
}

// Remove godoc
//
//	@summary	Remove a repository
//	@Security	ApiKeyAuth || cookies
//	@tags		Repo
//	@accept		json
//	@produce	json
//	@param		request	body		RemovePayload	true	"Repository url"
//	@success	200		{object}	app.GlobalResponse
//	@failure	400		{object}	app.GlobalResponse
//	@failure	500		{object}	app.GlobalResponse
//	@router		/repo [delete]
func Remove(c *fiber.Ctx) error {
	var payload RemovePayload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(app.GlobalResponse{
			Message: err.Error(),
		})
	}

	if payload.Repo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(app.GlobalResponse{
			Message: "missing repository url",
		})
	}

	if err := repository.Remove(payload.Repo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(app.GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(app.GlobalResponse{
		Message: "removed repository",
	})
}

// Update godoc
//
//	@summary	Update a repository
//	@tags		Repo
//	@Security	ApiKeyAuth || cookies
//	@accept		json
//	@produce	json
//	@param		request	body		PrivateRepoDetails	true	"Repository details"
//	@success	202		{object}	app.GlobalResponse
//	@failure	400		{object}	app.GlobalResponse
//	@failure	500		{object}	app.GlobalResponse
//	@router		/repo [put]
func Update(c *fiber.Ctx) error { // nolint:all
	var payload PrivateRepoDetails

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(app.GlobalResponse{
			Message: err.Error(),
		})
	}

	if (payload.URL == "" && payload.ImageRef == "") || payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(app.GlobalResponse{
			Message: "missing url, username or password in request body",
		})
	}

	if err := repository.Update(payload.URL, payload.ImageRef, payload.Username, payload.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(app.GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(app.GlobalResponse{
		Message: "Updated repository",
	})
}
