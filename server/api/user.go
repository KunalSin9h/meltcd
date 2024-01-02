package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core/auth"
)

// GetUsername godoc
//
//	@summary	Get username of current logged-in user
//	@tags		Users
//	@success	200	{string}	username
//	@failure	401
//	@router		/users/current [get]
func GetUsername(c *fiber.Ctx) error {
	// if login middleware will set this variable on login
	username := c.Locals("username").(string)
	if username == "" {
		return c.SendStatus(http.StatusUnauthorized)
	}

	return c.Status(http.StatusOK).SendString(username)
}

// GetUsers godoc
//
//	@summary	Get all the users
//	@tags		Users
//	@success	200	{object}	auth.AllUsers
//	@router		/users [get]
func GetUsers(c *fiber.Ctx) error {
	data := auth.GetAllUsers()

	return c.Status(200).JSON(data)
}

type ChangePasswordBody struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// ChangePassword godoc
//
//	@summary	Change password of user
//	@tags		Users
//	@accept		json
//	@param		request	body	ChangePasswordBody	true	"Change password body"
//	@success	200
//	@failure	400
//	@router		/users/{username} [patch]
func ChangePassword(c *fiber.Ctx) error {
	username := c.Params("username")

	var reqPay ChangePasswordBody

	if err := c.BodyParser(&reqPay); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	if reqPay.CurrentPassword == "" || reqPay.NewPassword == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	if done := auth.ChangePassword(username, reqPay.CurrentPassword, reqPay.NewPassword); !done {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.SendStatus(http.StatusOK)
}
