package daos

import (
	"github.com/eftakhairul/go-api-hack/cmd/models"
	"github.com/eftakhairul/go-api-hack/cmd/libs"
	"github.com/jinzhu/gorm"
)

// UserDAO persists user data in database
type UserDAO struct{
	DB *gorm.DB
}

// NewUserDAO creates a new UserDAO
func NewUserDAO(c libs.AppContext) *UserDAO {
	return &UserDAO{DB: c.DB}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *UserDAO) Get(id uint) (*models.User, error) {
	var user models.User

	// Query Database here...

	//user = models.User{
	//	Model: models.Model{ID: 1},
	//	FirstName: "Martin",
	//	LastName: "Heinz",
	//	Address: "Not gonna tell you",
	//	Email: "martin7.heinz@gmail.com"}

	// if using Gorm:
	err := dao.DB.Where("id = ?", id).
		First(&user).
		Error

	return &user, err
}
