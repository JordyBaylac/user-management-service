package api

import (
	"time"

	"github.com/JordyBaylac/user-management-service/users"
	"github.com/JordyBaylac/user-management-service/users/storage"
	"github.com/JordyBaylac/user-management-service/users/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup(idGenerator utils.UniqueIDGenerator) *fiber.App {
	app := fiber.New()

	// recover from panics
	app.Use(recover.New())

	// basic rate limiting
	app.Use(limiter.New(limiter.Config{
		Max:        30,
		Expiration: 1 * time.Minute,
	}))

	// setup routes and dependency injection
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
