package controllers

import (
	"fmt"

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/helper"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := claims["id"].(string)
	var user models.User
	database.DB.Db.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func UpdateMe(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	id := claims["id"].(string)
	role := claims["role"].(string)

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	if role != "admin" {
		user.Roles = role
	}

	result := database.DB.Db.Model(&models.User{}).Where("id = ?", id).Omit("password").Updates(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not update the user",
		})
	}

	return c.JSON(user)
}

func UpdateMePassword(c *fiber.Ctx) error {
	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id := claims["id"].(string)

	user := new(models.User)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := models.HashPassword(user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = hashedPassword
	result := database.DB.Db.Where("id = ?", id).Updates(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return c.SendString("User successfully updated")
}

func CreateUser(c *fiber.Ctx) error {

	claims, err := helper.GetClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	role := claims["role"].(string)
	if role == "admin" {

		fmt.Println(role)
	}

	fmt.Println(helper.GeneratePassword(8))
	// if role != "admin"{}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := models.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.ID, _ = models.GenerateUUID()

	user.Password = hashedPassword

	database.DB.Db.Create(&user)
	return c.Status(200).JSON(user)
}