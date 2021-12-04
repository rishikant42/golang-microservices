package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rishikant42/golang-microservices/mvc/services"
	"github.com/rishikant42/golang-microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	log.Printf("Process user id %v", userId)

	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user id must be integer",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiError)
		return
	}
	user, apiError := services.UserService.GetUser(userId)

	if apiError != nil {
		utils.RespondError(c, apiError)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
