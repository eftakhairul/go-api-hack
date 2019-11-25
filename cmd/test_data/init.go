package test_data

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/eftakhairul/go-api-hack/cmd/libs"
	"github.com/eftakhairul/go-api-hack/cmd/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

// Initializes application config and SQLite database used for testing
func Init() (*libs.AppConfig, *logrus.Logger, *gorm.DB) {
	// the test may be started from the home directory or a subdirectory
	// load application configurations
	conf, err := libs.LoadConfig("./config")
	if err != nil {
		panic(err)
	}

	appLog := libs.LoadAppLog()
	DB, DBErr := gorm.Open("sqlite3", ":memory:")
	DB.Exec("PRAGMA foreign_keys = ON") // SQLite defaults to `foreign_keys = off'`
	if DBErr != nil {
		panic(DBErr)
	}

	DB.AutoMigrate(&models.User{})
	return conf, appLog, DB
}

// Resets testing database - deletes all tables, creates new ones using GORM migration and populates them using `db.sql` file
func ResetDB(db *gorm.DB) {
	db.DropTableIfExists(&models.User{}) // Note: Order matters
	db.AutoMigrate(&models.User{})
	if err := runSQLFile(db, getSQLFile()); err != nil {
		panic(fmt.Errorf("error while initializing test database: %s", err))
	}
}

func getSQLFile() string {
	return "/test_data/db.sql" // on host use absolute path
}

func GetTestCaseFolder() string {
	return "/test_data/test_case_data" // on host use absolute path
}

// Executes SQL file specified by file argument
func runSQLFile(db *gorm.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if result := db.Exec(line); result.Error != nil {
			fmt.Println(line)
			return result.Error
		}
	}
	return nil
}
