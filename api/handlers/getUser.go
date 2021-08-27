package handlers

import (
	"github.com/JordyBaylac/user-management-service/users"
	"github.com/gofiber/fiber/v2"
)

type GetUserResponse struct {
	ID    string `json:"id" xml:"id,attr"`
	Email string `json:"email" xml:"email,attr"`
	Name  string `json:"name" xml:"name,attr"`
}

func HandleGetUser(service users.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var userID string
		if userID = c.Params("userID"); userID == "" {
			return c.Status(fiber.StatusBadRequest).SendString("userID param not found")
		}

		var result *users.User
		var err error
		if result, err = service.GetByID(userID); err != nil {
			return c.Status(fiber.StatusConflict).SendString(err.Error())
		}

		var response = GetUserResponse{
			ID:    result.ID,
			Email: result.Email,
			Name:  result.Name,
		}

		if err := c.Status(fiber.StatusCreated).JSON(response); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusExpectationFailed)
	}
}
