package apis

import "github.com/gin-gonic/gin"

// HealthGET returns health of application
func HealthGET(context *gin.Context) {
	context.JSON(200, gin.H{
		"status": "okay",
	})
}
