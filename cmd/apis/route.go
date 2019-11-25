package apis

import (
	"github.com/eftakhairul/go-api-hack/cmd/middlewares"
	models "github.com/eftakhairul/go-api-hack/cmd/models"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRoutes initiates all endpoints
func InitRoutes(httpEngine *gin.Engine) {

	//health check
	httpEngine.GET("/health", HealthGET)

	//Initate swagger
	httpEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := httpEngine.Group("/api/v1")
	{
		// Set Authorization middleware only for /v1 endpoints
		v1.Use(middlewares.AuthorizationMiddleware())
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", middlewares.PrepareBodyMiddleware(&models.User{}), PostUser)
	}
}
