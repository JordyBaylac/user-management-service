package api

import (
	"github.com/gofiber/fiber/v2"
)

const defaultPort = 8080

func Start() error {
	app := fiber.New()

	// routes
	setupRoutes(app, nil)

	return app.Listen(getServerAddress())
}
