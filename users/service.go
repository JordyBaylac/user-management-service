package users

import (
	"fmt"

	"github.com/JordyBaylac/user-management-service/users/domain"
	"github.com/JordyBaylac/user-management-service/users/storage"
)

type UserService interface {
	GetByID(userID string) (*domain.User, error)
	Create(email, name string) (*domain.User, error)
	Update(userID, newName string) (*domain.User, error)
}

type DefaultUserService struct {
	store storage.UserStorage
}

func NewUserService(storage storage.UserStorage) *DefaultUserService {
	return &DefaultUserService{storage}
}

func (service *DefaultUserService) Create(email, name string) (*domain.User, error) {
	store := service.store
	if exists := store.ExistByEmail(email); exists {
		return nil, fmt.Errorf("user with email %s is already present", email)
	}

	var newUser *domain.User
	var err error

	if newUser, err = store.CreateUser(email, name); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (service *DefaultUserService) GetByID(userID string) (*domain.User, error) {
	store := service.store
	var user *domain.User

	if user = store.GetByID(userID); user == nil {
		return nil, fmt.Errorf("user with id %s do not exist", userID)
	}

	return user, nil
}

func (service *DefaultUserService) Update(userID, newName string) (*domain.User, error) {
	var existingUser *domain.User
	var err error
	if existingUser, err = service.GetByID(userID); err != nil {
		return nil, fmt.Errorf("user with id %s do not exist", userID)
	}

	// update name only
	existingUser.Name = newName

	return existingUser, nil
}
