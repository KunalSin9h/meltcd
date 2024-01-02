package auth

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2/log"
	authPass "github.com/meltred/meltcd/internal/core/auth/password"
)

type User struct {
	Username     string    `json:"username"`
	PasswordHash string    `json:"passwordHash"` // hash passwords
	Role         UserRole  `json:"rol"`
	LastLoggedIn time.Time `json:"lastLoggedIn"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserRole string

const (
	Admin   UserRole = "admin"
	General UserRole = "general"
)

var users []*User

// Data without password hash
func (u *User) getPublicData() User {
	t := *u
	t.PasswordHash = ""

	return t
}

type AllUsers struct {
	Data []User `json:"data"`
}

func GetAllUsers() AllUsers {
	var all []User

	for _, user := range users {
		all = append(all, user.getPublicData())
	}

	return AllUsers{
		Data: all,
	}
}

var argon2Param = authPass.Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func FindUser(username, password string) (bool, error) {
	for _, user := range users {
		if user.Username == username {
			match, err := authPass.ComparePasswordAndHash(password, user.PasswordHash)
			if err != nil {
				return false, err
			}

			return match, nil
		}
	}

	return false, nil
}

func InsertUser(username, password string, role UserRole) error {
	hash, err := authPass.GenerateFromPassword(password, &argon2Param)
	if err != nil {
		return err
	}

	user := User{
		Username:     username,
		PasswordHash: hash,
		Role:         role,
		CreatedAt:    time.Now(),
	}

	users = append(users, &user)
	return nil
}

func ChangePassword(username, currentPassword, newPassword string) bool {
	log.Info("Changing password for user", "username", username)
	for _, user := range users {
		if user.Username == username {
			log.Info("user found", "username", username)
			match, err := authPass.ComparePasswordAndHash(currentPassword, user.PasswordHash)
			if err != nil {
				return false
			}
			if !match {
				return false
			}

			newHash, err := authPass.GenerateFromPassword(newPassword, &argon2Param)
			if err != nil {
				return false
			}

			user.PasswordHash = newHash
			log.Info("Changed password", "username", username)
			return true
		}
	}

	return false
}

func UserLoginUpdateTime(username string) {
	for _, user := range users {
		if user.Username == username {
			user.LastLoggedIn = time.Now()
			break
		}
	}
}

func LoadUsers(data *[]byte) error {
	return json.Unmarshal(*data, &users)
}

func GetUsers() (*[]byte, error) {
	data, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
