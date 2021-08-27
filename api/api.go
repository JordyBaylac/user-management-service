package api

import (
	"github.com/JordyBaylac/user-management-service/users"
	"github.com/JordyBaylac/user-management-service/users/storage"
	"github.com/JordyBaylac/user-management-service/users/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup() *fiber.App {
	app := fiber.New()

	// recover from panics
	app.Use(recover.New())

	// setup routes and dependency injection
	setupRoutes(app,
		users.NewUserService(
			storage.NewInMemoryStorage(utils.NewUUIDGenerator()),
		),
	)

	return app
}
