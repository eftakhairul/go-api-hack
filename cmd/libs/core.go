package libs

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type AppContext struct {
	Config *AppConfig
	Logger *logrus.Logger
	DB     *gorm.DB
}

func ginContextToAppContext(config *AppConfig, logger *logrus.Logger, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		appContext := &AppContext{config, logger, db}
		c.Set("appContext", appContext)
		c.Next()
	}
}

// LoadHTTPEngine initates gin framework and returns gin engine
func LoadHTTPEngine(config *AppConfig, logger *logrus.Logger, db *gorm.DB) *gin.Engine {
	// Creates a router without any middleware by default
	engine := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	engine.Use(gin.Logger())

	engine.Use(ginContextToAppContext(config, logger, db))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	engine.Use(gin.Recovery())

	return engine
}
