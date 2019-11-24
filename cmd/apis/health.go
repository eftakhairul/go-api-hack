package apis

import "github.com/gin-gonic/gin"

func HealthGET(context *gin.Context) {
	context.JSON(200, gin.H{
		"status": "okay",
	})
}
