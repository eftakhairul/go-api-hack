package apis

import "github.com/gin-gonic/gin"

// HealthGET returns health of application
func HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{"status": "okay"})
}
