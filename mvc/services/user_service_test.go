package services

import (
	"net/http"
	"testing"

	"github.com/rishikant42/golang-microservices/mvc/domain"
	"github.com/rishikant42/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDb(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "User 0 was not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.Id)
}
