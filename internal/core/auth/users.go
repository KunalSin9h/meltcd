package auth

import (
	"encoding/json"

	authPass "github.com/meltred/meltcd/internal/core/auth/password"
)

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"` // hash passwords
}

var users []*User

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

func InsertUser(username, password string) error {
	hash, err := authPass.GenerateFromPassword(password, &argon2Param)
	if err != nil {
		return err
	}

	user := User{
		Username:     username,
		PasswordHash: hash,
	}

	users = append(users, &user)
	return nil
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
