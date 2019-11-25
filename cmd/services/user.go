package services

import "github.com/eftakhairul/go-api-hack/cmd/models"

type userDAO interface {
	Get(id uint) (*models.User, error)
	Create(user *models.User) error
}

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

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (userService *UserService) Get(id uint) (*models.User, error) {
	return userService.dao.Get(id)
}
