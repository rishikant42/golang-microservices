package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

func RespondError(c *gin.Context, e *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(e.StatusCode, e)
		return
	}
	c.JSON(e.StatusCode, e)
}
