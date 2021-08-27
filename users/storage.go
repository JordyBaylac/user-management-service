package users

import "github.com/JordyBaylac/user-management-service/users/models"

type UserStorage interface {
	// queries
	ExistByEmail(email string) bool
	GetByID(userID string) *models.User

	// commands
	CreateUser(email, name string) (*models.User, error)
}
