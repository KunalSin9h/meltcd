package auth

import (
	"encoding/json"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"` // hash passwords
}

var users *[]User

func LoadUsers(data *[]byte) error {
	return json.Unmarshal(*data, &users)
}

func GetUsers() (*[]byte, error) {
	data, err := json.Marshal(&users)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
