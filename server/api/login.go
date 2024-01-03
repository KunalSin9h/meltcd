package api

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meltred/meltcd/internal/core/auth"
	"github.com/meltred/meltcd/internal/core/base58"
)

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login godoc
//
//	@summary	Login user
//	@tags		General
//	@accept		json
//	@param		request	body	LoginBody true "Login request body"
//	@success	200 string  string "Access Token"
//	@failure	400
//	@failure	500
//	@router		/login [post]
func Login(c *fiber.Ctx) error {
	var reqBody LoginBody

	if err := c.BodyParser(&reqBody); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	if reqBody.Username == "" || reqBody.Password == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	// check if username and password exits
	userExists, err := auth.FindUser(reqBody.Username, reqBody.Password)
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
	go auth.AddSession(token, reqBody.Username, expireTime)

	go auth.UserLoginUpdateTime(reqBody.Username)

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

func GenerateToken(n uint64) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
