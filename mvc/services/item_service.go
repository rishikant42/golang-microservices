package services

import (
	"net/http"

	"github.com/rishikant42/golang-microservices/mvc/domain"
	"github.com/rishikant42/golang-microservices/mvc/utils"
)

type itemService struct{}

var (
	ItemService itemService
)

func (i *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
