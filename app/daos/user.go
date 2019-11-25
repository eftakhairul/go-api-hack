package daos

import (
	"github.com/eftakhairul/go-api-hack/app/libs"
	"github.com/eftakhairul/go-api-hack/app/models"
	"github.com/jinzhu/gorm"
)

// UserDAO persists user data in database
type UserDAO struct {
	DB *gorm.DB
}

// NewUserDAO creates a new UserDAO
func NewUserDAO(c *libs.AppContext) *UserDAO {
	return &UserDAO{DB: c.DB}
}

// Find does the actual query to database, if user with specified id is not found error is returned
func (dao *UserDAO) Find(id uint) (*models.User, error) {
	var user models.User
	err := dao.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

// FindAll returns all the users from DB
func (dao *UserDAO) FindAll() (*[]models.User, error) {
	var users []models.User
	err := dao.DB.Find(&users).Error

	return &users, err
}

// Create create a new user into DB
func (dao *UserDAO) Create(user *models.User) error {
	// dao.DB.AutoMigrate(&models.User{})
	return dao.DB.Create(user).Error
}
