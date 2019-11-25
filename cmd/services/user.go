package services

import "github.com/eftakhairul/go-api-hack/cmd/models"

type userDAO interface {
	Find(id uint) (*models.User, error)
	FindAll() (*[]models.User, error)
	Create(user *models.User) error
}

// UserService Service
type UserService struct {
	dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Create create new user by userDAO
func (userService *UserService) Create(user *models.User) error {
	// All the business logic goes here
	//********************************
	// All the business logic goes here

	return userService.dao.Create(user)
}

// GetAll return all the users
func (userService *UserService) GetAll() (*[]models.User, error) {
	// All the business logic goes here
	//********************************
	// All the business logic goes here

	return userService.dao.FindAll()
}

// Get just retrieves one user using User DAO
func (userService *UserService) Get(id uint) (*models.User, error) {
	// All the business logic goes here
	//********************************
	// All the business logic goes here

	return userService.dao.Find(id)
}
