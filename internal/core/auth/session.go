package auth

import (
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

func SessionUpdateUsername(oldUsername, newUsername string) {
	for _, ses := range sessions {
		if ses.Username == oldUsername {
			ses.Username = newUsername
		}
	}
}

func VerifySession(token string) (string, bool) {
	for _, ses := range sessions {
		if ses.Token == token {
			if time.Now().After(ses.ExpireTime) {
				defer RemoveSession(token)

				return "", false
			}

			return ses.Username, true
		}
	}

	return "", false
}
