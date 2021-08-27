package handlers

import (
	"github.com/JordyBaylac/user-management-service/users"
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequest struct {
	Email string `json:"email" xml:"email,attr" validate:"required,email"`
	Name  string `json:"name" xml:"name,attr"`
}

type CreateUserResponse struct {
	ID    string `json:"id" xml:"id,attr"`
	Email string `json:"email" xml:"email,attr"`
	Name  string `json:"name" xml:"name,attr"`
}

func HandleCreateUser(service users.UserService) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		req := new(CreateUserRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		errors := ValidateStruct(req)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}

		var result *users.User
		var err error
		if result, err = service.Create(req.Email, req.Name); err != nil {
			return c.Status(fiber.StatusConflict).SendString(err.Error())
		}

		var response = CreateUserResponse{
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
