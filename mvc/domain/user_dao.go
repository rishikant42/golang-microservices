package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rishikant42/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "RK", LastName: "Sharma", Email: "rshkntshrm@gmail.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("we are accessing the DB")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "Not found",
	}
}
