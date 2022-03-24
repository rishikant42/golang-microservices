package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	fmt.Println("Starting app")
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func mapUrls() {
	router.GET("/users/:user_id", GetUser)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("user-id: %d", userId))
	}

	log.Printf("Process user id %v", userId)

	c.JSON(http.StatusOK, fmt.Sprintf("user-id: %d", userId))
}

func main() {
	StartApp()
}
