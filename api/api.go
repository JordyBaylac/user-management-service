package api

import (
	"github.com/JordyBaylac/user-management-service/users"
	"github.com/JordyBaylac/user-management-service/users/storage"
	"github.com/JordyBaylac/user-management-service/users/utils"
	"github.com/gofiber/fiber/v2"
)

func Setup(idGenerator utils.UniqueIDGenerator) *fiber.App {
	app := fiber.New()

	setupMiddlewares(app)

	// setup dependency injection
	generator := idGenerator
	if idGenerator == nil {
		generator = utils.NewUUIDGenerator()
	}

	setupRoutes(app,
		users.NewUserService(
			storage.NewInMemoryStorage(generator),
		),
	)

	return app
}
