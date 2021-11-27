package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, "Not found", err.Code)
	assert.Equal(t, "User 0 was not found", err.Message)
}
func TestGetUserFound(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, uint64(123), user.Id)
	assert.Equal(t, "RK", user.FirstName)
	assert.Equal(t, "Sharma", user.LastName)
	assert.EqualValues(t, "rshkntshrm@gmail.com", user.Email)
}
