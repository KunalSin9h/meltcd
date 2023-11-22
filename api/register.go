package api

import (
	"meltred/meltcd/internal/core"
	"meltred/meltcd/internal/core/application"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var app application.Application

	if err := c.BodyParser(&app); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if err := core.Register(&app); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return c.SendStatus(http.StatusAccepted)
}
