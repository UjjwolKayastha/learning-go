package services

import (
	"github.com/jinzhu/copier"
	"github.com/ujjwolkayastha/learn-clean-gin/api/repositories"
	"github.com/ujjwolkayastha/learn-clean-gin/lib"
	"github.com/ujjwolkayastha/learn-clean-gin/models"
)

// UserService service layer
type UserService struct {
	logger     lib.Logger
	repository repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(logger lib.Logger, repository repositories.UserRepository) UserService {
	return UserService{
		logger:     logger,
		repository: repository,
	}
}

// GetAllUser get all the user
func (s UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.repository.GetAll()
	return users, err
}

// GetOneUser gets one user
func (s UserService) GetOneUser(id uint) (models.User, error) {
	user, err := s.repository.GetOne(id)
	return user, err
}

// CreateUser call to create the user
func (s UserService) CreateUser(user models.User) error {
	_, err := s.repository.Save(user)
	return err
}

// UpdateUser updates the user
func (s UserService) UpdateUser(id uint, user models.User) error {

	userDB, err := s.GetOneUser(id)
	if err != nil {
		return err
	}

	copier.Copy(&userDB, &user)

	userDB.ID = id

	_, err = s.repository.Update(userDB)
	return err
}

// DeleteUser deletes the user
func (s UserService) DeleteUser(id uint) error {
	return s.repository.Delete(id)
}
