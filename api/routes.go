package api

import (
	"github.com/JordyBaylac/user-management-service/api/handlers"
	"github.com/JordyBaylac/user-management-service/users"
	"github.com/gofiber/fiber/v2"
)

// setupRoutes will define api routes
func setupRoutes(
	app *fiber.App,
	userService users.UserService,
) {
	app.Group("/users").
		Post("/", handlers.HandleCreateUser(userService)).
		Get("/:id", handlers.HandleGetUser(userService)).
		Patch("/:id", handlers.HandleUpdateUser(userService))
}
