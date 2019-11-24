package libs

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type AppContext struct {
	*gin.Context
	Logger *logrus.Logger
	DB     *gorm.DB
}

// LoadHTTPEngine initates gin framework and returns gin engine
func LoadHTTPEngine(logger *logrus.Logger) *gin.Engine {
	// Creates a router without any middleware by default
	engine := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	engine.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	engine.Use(gin.Recovery())

	return engine
}
