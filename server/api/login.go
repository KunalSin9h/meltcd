package api

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core/auth"
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
	userExists, err := auth.FindUser(username, password)
	if err != nil {
		return err
	}

	if !userExists {
		return c.SendStatus(http.StatusUnauthorized)
	}

	token, err := GenerateToken(32)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	expireTime := time.Now().Add(1 * time.Hour)
	go auth.AddSession(token, username, expireTime)

	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    token,
		Expires:  expireTime,
		Secure:   true,
		SameSite: "Strict",
		HTTPOnly: true,
	})

	return c.Redirect("/")
}

func GenerateToken(n uint64) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
