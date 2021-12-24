package services

import (
	"net/http"
	"testing"

	"github.com/rishikant42/golang-microservices/src/api/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoCuncurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	output := make(chan repositories.CreateRepositoriesResults)

	services := repoService{}

	go services.createRepoCuncurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "Invalid repo name", result.Error.Message())
}
