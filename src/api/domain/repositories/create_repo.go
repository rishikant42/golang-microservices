package repositories

import (
	"strings"

	"github.com/rishikant42/golang-microservices/src/api/utils/errors"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateRepoRequest) Validate() errors.ApiError {
	r.Name = strings.TrimSpace(r.Name)

	if r.Name == "" {
		return errors.NewBadRequestError("Invalid repo name")
	}
	return nil
}

type CreateRepoResponse struct {
	Id    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int                         `json:"status"`
	Results    []CreateRepositoriesResults `json:"results"`
}

type CreateRepositoriesResults struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.ApiError     `json:"error"`
}
