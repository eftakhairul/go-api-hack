package main

import (
	"fmt"

	"github.com/eftakhairul/go-api-hack/cmd/apis"
	"github.com/eftakhairul/go-api-hack/cmd/config"
	"github.com/eftakhairul/go-api-hack/cmd/libs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_ "github.com/eftakhairul/go-api-hack/cmd/docs"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang API Hack
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email eftakhairul@gmail.com

// @license.name MIT
// @license.url https://github.com/eftakhairul/go-api-hack/blob/master/LICENSE

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}
	// config.Config.DB.AutoMigrate(&models.User{}) // This is needed for generation of schema for postgres image.
	defer config.Config.DB.Close()

	//Intiate Custom logger
	appLog := libs.LoadAppLog()

	//Initate Gin engine and load routes
	httpEngine := libs.LoadHTTPEngine(appLog)
	apis.InitRoutes(httpEngine)

	appLog.Info("Successfully connected to database")
	httpEngine.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
