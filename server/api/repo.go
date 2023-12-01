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

func RepoAdd(c *fiber.Ctx) error {
	var payload PrivateRepoDetails

	if err := c.BodyParser(&payload); err != nil {
		return err
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
		Message: "Added new repository",
	})
}

type RepoListData struct {
	Data []string `json:"data"`
}

func RepoList(c *fiber.Ctx) error {
	list := repository.List()

	return c.Status(fiber.StatusOK).JSON(RepoListData{
		Data: list,
	})
}
