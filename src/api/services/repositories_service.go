package services

import (
	"github.com/rishikant42/golang-microservices/src/api/config"
	"github.com/rishikant42/golang-microservices/src/api/domain/github"
	"github.com/rishikant42/golang-microservices/src/api/domain/repositories"
	"github.com/rishikant42/golang-microservices/src/api/provider/github_provider"
	"github.com/rishikant42/golang-microservices/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	// input.Name = strings.Trim(input.Name)

	if input.Name == "" {
		return nil, errors.NewBadRequestError("Invalid repo name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}
