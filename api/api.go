package api

import (
	"github.com/JordyBaylac/user-management-service/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const defaultPort = 8080

func Start() error {
	app := fiber.New()

	// recover from panics
	app.Use(recover.New())

	setupRoutes(app, users.CreateInMemoryUserService())
	return app.Listen(getServerAddress())
}
