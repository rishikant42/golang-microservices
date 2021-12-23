package app

import (
	"github.com/rishikant42/golang-microservices/src/api/controllers/health"
	"github.com/rishikant42/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/health", health.HealthCheck)
	router.POST("/repositories", repositories.CreateRepo)
}
