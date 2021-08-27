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
	users := app.Group("/users")
	users.Post("/", handlers.HandleCreateUser(userService))
	users.Get("/:userID", handlers.HandleGetUser(userService))
	users.Patch("/:userID", handlers.HandleUpdateUser(userService))
}
