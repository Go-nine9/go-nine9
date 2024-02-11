package controllers

import (
	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PrestationRequest struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	ServiceID string  `json:"service_id"`
}

func CreatePrestation(c *fiber.Ctx) error {
	var prestationRequest PrestationRequest
	if err := c.BodyParser(&prestationRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	serviceID, err := uuid.Parse(prestationRequest.ServiceID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid service ID",
		})
	}

	prestation := models.Prestation{
		ID:        uuid.New(),
		Name:      prestationRequest.Name,
		Price:     prestationRequest.Price,
		ServiceID: &serviceID,
	}

	result := database.DB.Db.Create(&prestation)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	response := fiber.Map{
		"message":    "Prestation créée avec succès",
		"prestation": prestation,
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}
