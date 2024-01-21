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

	result := database.DB.Db.Create(&salon)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
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
	result = database.DB.Db.First(&manager, "id = ?", managerID)

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

// Get all salons with their staff associated
func GetSalons(c *fiber.Ctx) error {
	var salons []models.Salon
	result := database.DB.Db.Preload("Users").Find(&salons)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.JSON(salons)
}

// Retrieve the salon by ID with staff associated
func GetSalonById(c *fiber.Ctx) error {
	id := c.Params("id")
	var salon models.Salon
	result := database.DB.Db.Preload("Users").Where("id = ?", id).First(&salon)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}
	return c.JSON(salon)
}

// Update salon data
func UpdateSalon(c *fiber.Ctx) error {
	id := c.Params("id")
	salon := new(models.Salon)

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

// Add staff
func AddStaff(c *fiber.Ctx) error {
	id := c.Params("id")
	var users []models.User
	if err := c.BodyParser(&users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Convert the Salon Id string into UUID
	salonId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid salon ID format",
		})
	}

	// Create user for each iteration of array
	for i := 0; i < len(users); i++ {
		var user models.User
		user = users[i]
		// Assign the salonId to staff
		user.SalonID = &salonId
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

	return c.SendString("Salon successfully updated")
}
