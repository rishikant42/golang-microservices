package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func StartApp() {
	fmt.Println("Statring app on 8080 port")
	router = gin.Default()

	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
