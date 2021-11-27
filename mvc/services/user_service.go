package services

import (
	"github.com/rishikant42/golang-microservices/mvc/domain"
	"github.com/rishikant42/golang-microservices/mvc/utils"
)

type userService struct {
}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
