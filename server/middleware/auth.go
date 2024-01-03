package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core/auth"
)

func VerifyUser(c *fiber.Ctx) error {
	authToken := c.Cookies("authToken", "")
	if authToken == "" {
		return c.Status(http.StatusUnauthorized).SendString("missing authentication token, login first")
	}

	username, tokenExists := auth.VerifySession(authToken)
	if !tokenExists {
		return c.Status(http.StatusUnauthorized).SendString("missing authentication token, login first")
	}

	c.Locals("username", username)
	return c.Next()
}
