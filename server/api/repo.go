package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core/repository"
)

type PrivateRepoDetails struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type GlobalResponse struct {
	Message string `json:"message"`
}

// RepoAdd godoc
//
//	@summary	Add a new repository
//	@tags		Repo
//	@accept		json
//	@produce	json
//	@param		request	body		PrivateRepoDetails	true	"Repository details"
//	@success	202		{object}	GlobalResponse
//	@failure	400		{object}	GlobalResponse
//	@failure	500		{object}	GlobalResponse
//	@router		/repo [post]
func RepoAdd(c *fiber.Ctx) error { // nolint:all
	var payload PrivateRepoDetails

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	if payload.URL == "" || payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: "missing url, username or password in request body",
		})
	}

	url, _ := strings.CutSuffix(payload.URL, "/")

	if err := repository.Add(url, payload.Username, payload.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(GlobalResponse{
		Message: "Added Repository",
	})
}

type RepoListData struct {
	Data []string `json:"data"`
}

// RepoList godoc
//
//	@summary	Get a list all repositories
//	@tags		Repo
//	@produce	json
//	@success	200	{object}	RepoListData
//	@router		/repo [get]
func RepoList(c *fiber.Ctx) error {
	list := repository.List()

	return c.Status(fiber.StatusOK).JSON(RepoListData{
		Data: list,
	})
}

type RepoRemovePayload struct {
	Repo string `json:"repo"`
}

// RepoRemove godoc
//
//	@summary	Remove a repository
//	@tags		Repo
//	@accept		json
//	@produce	json
//	@param		request	body		RepoRemovePayload	true	"Repository url"
//	@success	200		{object}	GlobalResponse
//	@failure	400		{object}	GlobalResponse
//	@failure	500		{object}	GlobalResponse
//	@router		/repo [delete]
func RepoRemove(c *fiber.Ctx) error {
	var payload RepoRemovePayload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	if payload.Repo == "" {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: "missing repository url",
		})
	}

	if err := repository.Remove(payload.Repo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(GlobalResponse{
		Message: "removed repository",
	})
}

// RepoUpdate godoc
//
//	@summary	Update a repository
//	@tags		Repo
//	@accept		json
//	@produce	json
//	@param		request	body		PrivateRepoDetails	true	"Repository details"
//	@success	202		{object}	GlobalResponse
//	@failure	400		{object}	GlobalResponse
//	@failure	500		{object}	GlobalResponse
//	@router		/repo [put]
func RepoUpdate(c *fiber.Ctx) error { // nolint:all
	var payload PrivateRepoDetails

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	if payload.URL == "" || payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(GlobalResponse{
			Message: "missing url, username or password in request body",
		})
	}

	url, _ := strings.CutSuffix(payload.URL, "/")

	if err := repository.Update(url, payload.Username, payload.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(GlobalResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(GlobalResponse{
		Message: "Updated repository",
	})
}
