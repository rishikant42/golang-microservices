package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/rishikant42/golang-microservices/mvc/services"
	"github.com/rishikant42/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	log.Printf("Process user id %v", userId)

	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user id must be integer",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiError := services.UserService.GetUser(userId)

	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
