package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core/auth"
)

func VerifyUser(c *fiber.Ctx) error {
	// See if authToken (Access Token) is in Cookies
	authToken := c.Cookies("authToken", "")

	// see if token is in Authorization: Bearer <token> Header
	authHeader := c.Request().Header.Peek("Authorization")

	bearerToken, _ := strings.CutPrefix(string(authHeader), "Bearer ")

	// BearerToken is same as authToken but extracted from different place,
	// authToken extracted from cookies are most probably coming from browser ui
	// bearerToken is coming from CLI

	var token string
	if authToken != "" {
		token = authToken
	}
	if bearerToken != "" {
		token = bearerToken
	}

	username, sessionExists := auth.VerifySession(token)

	if token == "" || !sessionExists {
		return c.Status(http.StatusUnauthorized).SendString("missing authentication token in cookies\nOR session is expired\nOR you just restart the app\n→ login again!")
	}

	c.Locals("username", username)
	return c.Next()
}
