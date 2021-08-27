package storage

import (
	"sync"

	"github.com/JordyBaylac/user-management-service/users/models"
	"github.com/JordyBaylac/user-management-service/users/utils"
)

type InMemoryStorage struct {
	generator utils.UniqueIDGenerator

	// map of user by ID
	users map[string]*models.User

	// set to identify existing emails
	emails map[string]bool

	// lock to ensure safe access and avoid data races
	lock sync.RWMutex
}

func NewInMemoryStorage(generator utils.UniqueIDGenerator) *InMemoryStorage {
	return &InMemoryStorage{
		generator: generator,
		users:     make(map[string]*models.User),
		emails:    make(map[string]bool),
	}
}

func (storage *InMemoryStorage) ExistByEmail(email string) bool {
	storage.lock.RLock()
	defer storage.lock.RUnlock()

	if exists := storage.emails[email]; exists {
		return true
	}

	return false
}

func (storage *InMemoryStorage) GetByID(userID string) *models.User {
	var user *models.User
	var found bool

	storage.lock.RLock()
	defer storage.lock.RUnlock()

	if user, found = storage.users[userID]; !found {
		return nil
	}

	return user
}

func (storage *InMemoryStorage) CreateUser(email, name string) (*models.User, error) {
	uniqueID := storage.generator.GenerateID()
	newUser := &models.User{
		ID:    uniqueID,
		Email: email,
		Name:  name,
	}

	storage.lock.Lock()
	defer storage.lock.Unlock()
	storage.emails[email] = true
	storage.users[uniqueID] = newUser

	return newUser, nil
}
