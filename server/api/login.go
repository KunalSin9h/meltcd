package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Login godoc
//
//	@summary	Login user
//	@tags		General
//	@success	302
//	@failure	400
//	@failure	401
//	@router		/login [post]
func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	// check if username and password exits
	// if not
	// return http.StatusUnauthorized
	// if exist create a token and store it in cookie
	// redirect to /

	return nil
}
