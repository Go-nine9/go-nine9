package controllers

import (
	"net/http"

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SalonRequest struct {
	User    []models.User `json:"user"`
	Salon   models.Salon  `json:"salon"`
	Manager string        `json:"manager"`
}

type SalonResponse struct {
	Salon models.Salon `json:"salon"`
}

func CreateSalon(c *fiber.Ctx) error {
	var salonRequest SalonRequest
	if err := c.BodyParser(&salonRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	salon := models.Salon{
		Name:    salonRequest.Salon.Name,
		Address: salonRequest.Salon.Address,
		Phone:   salonRequest.Salon.Phone,
	}

	salon.ID, _ = models.GenerateUUID()

	database.DB.Db.Create(&salon)

	// Check if the creation was successful
	if database.DB.Db.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": database.DB.Db.Error.Error(),
		})
	}

	// Convert the manager ID from string to UUID
	managerID, err := uuid.Parse(salonRequest.Manager)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid manager ID format",
		})
	}

	// Find the manager by ID
	var manager models.User
	result := database.DB.Db.First(&manager, "id = ?", managerID)

	// Check if there is an error and the manager doesn't exist
	if result.Error != nil && result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Manager not found",
		})
	}

	// Associate the salon to manager or create a new manager if not found
	manager.SalonID = &salon.ID
	database.DB.Db.Save(&manager)

	// Create staff users
	users := salonRequest.User
	for i := 0; i < len(users); i++ {
		var user models.User
		user = users[i]
		// Assign the salonId to staff
		user.SalonID = &salon.ID
		user.Roles = "staff"
		user.ID, _ = models.GenerateUUID()

		hashedPassword, err := models.HashPassword(user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to hash password",
			})
		}
		user.Password = hashedPassword

		database.DB.Db.Create(&user)

	}

	return c.Status(http.StatusOK).JSON("Le salon a bien été crée")
}
