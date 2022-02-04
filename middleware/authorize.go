package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
