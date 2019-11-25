package main

import (
	"fmt"

	"github.com/eftakhairul/go-api-hack/cmd/apis"
	_ "github.com/eftakhairul/go-api-hack/cmd/docs"
	"github.com/eftakhairul/go-api-hack/cmd/libs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	conf, err := libs.LoadConfig("./config")
	if err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	appLog := libs.LoadAppLog()

	//DB, dberr = gorm.Open("postgres", config.Config.DSN)
	DB, dberr := gorm.Open("sqlite3", "test.db")
	if dberr != nil {
		panic(dberr)
	}

	//Intiate Custom logger
	appLog.Info("Successfully connected to database")

	// config.Config.DB.AutoMigrate(&models.User{}) // This is needed for generation of schema for postgres image.
	defer DB.Close()

	//Initate Gin engine and load routes
	httpEngine := libs.LoadHTTPEngine(conf, appLog, DB)
	apis.InitRoutes(httpEngine)

	httpEngine.Run(fmt.Sprintf("%v:%v", conf.Hostname, conf.Port))
}
