package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitializeMiddlewares(router *gin.Engine) {
	router.Use(CORS)
}
