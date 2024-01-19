package controller

import (
	"github.com/go-nine9/go-nine9/models"
	"github.com/go-nine9/go-nine9/services"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"

func sendSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Create a user with the provided data
	newUser := models.User{
		Firstname: data["firstname"],
		Lastname:  data["lastname"],
		Email:     data["email"],
		Password:  data["password"],
		Roles:     "user",
	}
	if err := services.RegisterUser(newUser); err != nil {
		return err
	}
	token, err := services.GenerateJWT(newUser)
	if err != nil {
		return err
	}

	if roles, ok := c.Locals("userRoles").(string); ok && roles == "admin" {
		if err := services.AddUserByAdmin(newUser); err != nil {
			return err
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"jwt":       token,
		"email":     newUser.Email,
		"firstname": newUser.Firstname,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	email := data["email"]
	password := data["password"]

	user, err := services.FindUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := services.VerifyPassword(string([]byte(user.Password)), password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	token, err := services.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"jwt":       token,
		"email":     user.Email,
		"firstname": user.Firstname,
	})
}
