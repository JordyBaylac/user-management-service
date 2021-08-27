package users

import (
	uuid "github.com/satori/go.uuid"
)

func generateUniqueID() string {
	return uuid.NewV4().String()
}
