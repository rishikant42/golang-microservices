package app

import (
	"github.com/gin-gonic/gin"
	// log_option_a "github.com/rishikant42/golang-microservices/src/api/log/option_a"
	log_option_b "github.com/rishikant42/golang-microservices/src/api/log/option_b"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	// log_option_a.Info("About to map urls", "step:1", "status:pending")
	log_option_b.Info("About to map urls", log_option_b.Field("step", "1"), log_option_b.Field("status", "pending"))
	mapUrls()
	// log_option_a.Info("Urls successfully mapped", "step:2", "status:success")
	log_option_b.Info("Urls successfully mapped", log_option_b.Field("step", "2"), log_option_b.Field("status", "success"))

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
