package oauth

import (
	"fmt"

	"github.com/rishikant42/golang-microservices/src/api/utils/errors"
)

const (
	queryGetUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"fede": {Id: 123, Username: "fede"},
		"rk":   {Id: 454, Username: "rishi"},
	}
)

func GetUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]

	if user == nil {
		return nil, errors.NewNotFoundApiError(fmt.Sprintf("no user with username: %s", username))
	}
	return user, nil
}
