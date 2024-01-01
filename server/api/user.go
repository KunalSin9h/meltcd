package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetUser godoc
//
//	@summary	Get username of current logged-in user
//	@tags		User
//	@success	200 {string} username
//	@failure	401
//	@router		/user [get]
func GetUser(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	if username == "" {
		return c.SendStatus(http.StatusUnauthorized)
	}

	return c.Status(http.StatusOK).SendString(username)
}
