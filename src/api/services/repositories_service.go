package services

import (
	"net/http"
	"sync"

	"github.com/rishikant42/golang-microservices/src/api/config"
	"github.com/rishikant42/golang-microservices/src/api/domain/github"
	"github.com/rishikant42/golang-microservices/src/api/domain/repositories"
	"github.com/rishikant42/golang-microservices/src/api/provider/github_provider"
	"github.com/rishikant42/golang-microservices/src/api/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
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
func (s *repoService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResults)
	output := make(chan repositories.CreateReposResponse)

	defer close(output)

	var wg sync.WaitGroup

	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoCuncurrent(current, input)
	}
	wg.Wait()
	close(input)

	result := <-output

	successCreations := 0

	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResults, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResults{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *repoService) createRepoCuncurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResults) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResults{Error: err}
		return
	}
	result, err := s.CreateRepo(input)

	if err != nil {
		output <- repositories.CreateRepositoriesResults{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResults{Response: result}

	// request := github.CreateRepoRequest{
	// 	Name:        input.Name,
	// 	Description: input.Description,
	// 	Private:     false,
	// }

	// result, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	// if err != nil {
	// 	output <- repositories.CreateRepositoriesResults{
	// 		Error: errors.NewApiError(err.StatusCode, err.Message),
	// 	}
	// 	return
	// }

	// output <- repositories.CreateRepositoriesResults{
	// 	Response: &repositories.CreateRepoResponse{
	// 		Id:    result.Id,
	// 		Name:  result.Name,
	// 		Owner: result.Owner.Login,
	// 	},
	// }
}
