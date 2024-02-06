package controllers

import (

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetMySalons(c *fiber.Ctx) error {
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	if claims["salonID"] == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "noon",
		})
	}

	salonId := claims["salonID"].(string)

	var salons []models.Salon
	result := database.DB.Db.
		Preload("User").
		Preload("User.Slots").
		Find(&salons, "id = ?", salonId)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(salons)
}

func GetSalons(c *fiber.Ctx) error {
	var salons []models.Salon
	result := database.DB.Db.
		Preload("User").
		Preload("User.Salon").
		Preload("User.Slots").
		Find(&salons)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(salons)
}

func CreateSalon(c *fiber.Ctx) error {
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	salon := new(models.Salon)
	if err := c.BodyParser(salon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	salon.ID, _ = models.GenerateUUID()
	if err := c.BodyParser(salon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	role := claims["role"].(string)
	if role == "user" {
		var user models.User
		userID := claims["id"].(string)
		user.Roles = "manager"
		user.SalonID = &salon.ID
		result := database.DB.Db.Where("id = ?", userID).Updates(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
	}

	if role == "manager" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "You have already a salon",
		})
	}

	for _, user := range salon.User {
		user.SalonID = &salon.ID
		user.Roles = "employee"
		user.ID, _ = models.GenerateUUID()
		hashedPassword, err := models.HashPassword(user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to hash password",
			})
		}
		user.Password = hashedPassword
		result := database.DB.Db.Create(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
	}

	result := database.DB.Db.Create(&salon)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.Status(200).JSON("salon")
}

func GetSalonById(c *fiber.Ctx) error {
	id := c.Params("id")
	var salon models.Salon
	result := database.DB.Db.Preload("User").Preload("User.Slots").Where("id = ?", id).First(&salon)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.JSON(salon)
}

func AddStaff(c *fiber.Ctx) error {
	salon := new(models.Salon)
	if err := c.BodyParser(salon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, user := range salon.User {
		user.SalonID = &salon.ID
		user.Roles = "employee"
		user.ID, _ = models.GenerateUUID()
		hashedPassword, err := models.HashPassword(user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to hash password",
			})
		}
		user.Password = hashedPassword
		result := database.DB.Db.Create(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": result.Error.Error(),
			})
		}
	}
	return c.Status(200).JSON(salon)
}

func UpdateSalon(c *fiber.Ctx) error {
	id := c.Params("id")
	salon := new(models.Salon)
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})

	}
	role := claims["role"].(string)
	if role == "manager" {
		id = claims["salonID"].(string)
	}

	if err := c.BodyParser(salon); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	result := database.DB.Db.Where("id = ?", id).Updates(&salon)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update salon",
		})
	}
	return c.SendString("Salon successfully updated")
}

func DeleteSalon(c *fiber.Ctx) error {
	id := c.Params("id")
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	role := claims["role"].(string)
	if role == "manager" {
		id = claims["salonID"].(string)
	}

	var salon models.Salon
	result := database.DB.Db.Where("id = ?", id).Delete(&salon)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}
	return c.Status(200).JSON(salon)
}
