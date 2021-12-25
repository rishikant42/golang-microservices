package oauth

import (
	"fmt"

	"github.com/rishikant42/golang-microservices/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	at.Token = fmt.Sprintf("USER_%d", at.UserId)
	tokens[at.Token] = at
	return nil
}

func GetAccessTokenByToken(token string) (*AccessToken, errors.ApiError) {
	accessToken := tokens[token]
	if accessToken == nil {
		return nil, errors.NewNotFoundApiError("no access token found")
	}
	return accessToken, nil
}
