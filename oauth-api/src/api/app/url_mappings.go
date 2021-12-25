package app

import (
	"github.com/rishikant42/golang-microservices/oauth-api/src/api/controllers/oauth"
	"github.com/rishikant42/golang-microservices/src/api/controllers/health"
)

func mapUrls() {
	router.GET("/health", health.HealthCheck)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
