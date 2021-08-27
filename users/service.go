package users

import "fmt"

type UserService interface {
	Create(email string, name string) (*User, error)
	GetByID(userID string) (*User, error)
	Update(userID string, newName string) (*User, error)
}

type InMemoryUserService struct {
	// memory storage for storing users by ID
	users map[string]*User

	// set to identify existing emails
	emails map[string]bool
}

func CreateInMemoryUserService() InMemoryUserService {
	return InMemoryUserService{
		users:  make(map[string]*User),
		emails: make(map[string]bool),
	}
}

func (service *InMemoryUserService) Create(email string, name string) (*User, error) {
	if exists := service.emails[email]; exists {
		return nil, fmt.Errorf("user with email %s is already present", email)
	}

	uniqueID := generateUniqueID()
	user := User{
		ID:    uniqueID,
		Email: email,
		Name:  name,
	}

	service.emails[email] = true
	service.users[uniqueID] = &user

	return &user, nil
}

func (service *InMemoryUserService) GetByID(userID string) (*User, error) {
	if _, found := service.users[userID]; !found {
		return nil, fmt.Errorf("user with id %s is not present", userID)
	}

	return service.users[userID], nil
}

func (service *InMemoryUserService) Update(userID string, newName string) (*User, error) {
	if _, found := service.users[userID]; !found {
		return nil, fmt.Errorf("user with id %s is not present", userID)
	}

	existingUser := service.users[userID]
	existingUser.Name = newName

	return existingUser, nil
}
