package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core/auth"
	"github.com/meltred/meltcd/internal/core/base58"
)

// Login godoc
//
//	@summary	Login user
//	@tags		General
//	@security	BasicAuth
//	@success	200	string	string
//	@failure	400
//	@failure	401
//	@failure	500
//	@router		/login [post]
func Login(c *fiber.Ctx) error {
	username, password := extractFromBasic(c)

	if username == "" || password == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	// check if username and password exits
	userExists, err := auth.FindUser(username, password)
	if err != nil {
		return err
	}

	if !userExists {
		return c.SendStatus(http.StatusUnauthorized)
	}

	token, err := base58.New(16)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	token = fmt.Sprintf("api_%s", token)

	expireTime := time.Now().Add(1 * time.Hour)
	go auth.AddSession(token, username, expireTime)

	go auth.UserLoginUpdateTime(username)

	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    token,
		Expires:  expireTime,
		Secure:   true,
		SameSite: "Strict",
		HTTPOnly: true,
	})

	return c.Status(http.StatusOK).SendString(token)
}

func extractFromBasic(c *fiber.Ctx) (string, string) {
	value := string(c.Request().Header.Peek("Authorization"))

	token, _ := strings.CutPrefix(value, "Basic ")

	creds, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", ""
	}

	credentials := string(creds)

	userPass := strings.SplitN(credentials, ":", 2)
	if len(userPass) != 2 {
		return "", ""
	}

	return userPass[0], userPass[1]
}
