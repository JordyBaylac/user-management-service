package storage

import "github.com/JordyBaylac/user-management-service/users/domain"

type UserStorage interface {
	// queries
	ExistByEmail(email string) bool
	GetByID(userID string) *domain.User

	// commands
	CreateUser(email, name string) (*domain.User, error)
}
