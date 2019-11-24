package middlewares

import "github.com/gin-gonic/gin"

func PrepareBodyMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
}
