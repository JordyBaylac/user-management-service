package handlers

import (
	"github.com/JordyBaylac/user-management-service/api/validations"
	"github.com/JordyBaylac/user-management-service/users"
	"github.com/JordyBaylac/user-management-service/users/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateUserRequest struct {
	Name string `json:"name" xml:"name" validate:"required"`
}

type UpdateUserResponse struct {
	ID    string `json:"id" xml:"id"`
	Email string `json:"email" xml:"email"`
	Name  string `json:"name" xml:"name"`
}

func HandleUpdateUser(service users.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// parse request
		var userID string
		if userID = c.Params("userID"); userID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorMessage": "userID param not found",
			})
		}
		req := new(UpdateUserRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorMessage": err.Error(),
			})
		}

		// validate request
		errors := validations.ValidateStruct(req)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errorMessage": "validation error",
				"details":      errors,
			})
		}

		// call business service
		var result *models.User
		var err error
		if result, err = service.Update(userID, req.Name); err != nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"errorMessage": err.Error(),
			})
		}

		// reply
		var response = UpdateUserResponse{
			ID:    result.ID,
			Email: result.Email,
			Name:  result.Name,
		}
		if err := c.Status(fiber.StatusOK).JSON(response); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errorMessage": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
