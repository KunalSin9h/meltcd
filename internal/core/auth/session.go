package auth

import (
	"fmt"
	"time"
)

type Session struct {
	Token      string    `json:"authToken"`
	Username   string    `json:"username"`
	ExpireTime time.Time `json:"expireTime"`
}

var sessions []*Session

func AddSession(token, username string, expireTime time.Time) {
	sessions = append(sessions, &Session{
		Token:      token,
		Username:   username,
		ExpireTime: expireTime,
	})

	fmt.Println(sessions)
}

func RemoveSession(token string) {
	var result []*Session

	for _, ses := range sessions {
		if ses.Token != token {
			result = append(result, ses)
		}
	}

	sessions = result
}
