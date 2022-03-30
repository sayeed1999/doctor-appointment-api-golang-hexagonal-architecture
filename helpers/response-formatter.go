package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, message string, body interface{}) {
	if message == "" {
		message = http.StatusText(code)
	}
	c.JSON(
		code,
		gin.H{
			"code":       code,
			"statusText": http.StatusText(code),
			"message":    message,
			"body":       body,
		},
	)
}
