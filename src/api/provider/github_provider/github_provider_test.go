package github_provider

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/rishikant42/golang-microservices/src/api/clients/restclient"
	"github.com/rishikant42/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMock()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := GetAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	// restclient.StartMock()
	restclient.FlushMocks()
	restclient.AddMock(restclient.Mock{
		Url:        urlCreateRepo,
		HttpMethod: http.MethodPost,
		Response:   nil,
		Err:        errors.New("Invalid restclient response"),
	})
	resp, err := CreateRepo("", github.CreateRepoRequest{})
	restclient.StopMock()
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Invalid restclient response", err.Message)
}
