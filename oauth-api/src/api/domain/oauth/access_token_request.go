package oauth

import (
	"strings"

	"github.com/rishikant42/golang-microservices/src/api/utils/errors"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AccessTokenRequest) Validate() errors.ApiError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}
	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
