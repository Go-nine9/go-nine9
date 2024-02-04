package controllers

import (
	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
)

func GetReservationById(c *fiber.Ctx) error {
	slotId := c.Params("slotId")
	var reservation []models.Reservation
	result := database.DB.Db.
		Preload("User").
		Preload("Slot").
		Preload("Slot.HairdressingStaff").
		Preload("Slot.HairdressingStaff.Salon").
		Where("reservations.slot_id = ?", slotId).
		Find(&reservation)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(reservation)
}

func CreateReservation(c *fiber.Ctx) error {
	reservation := new(models.Reservation)
	if err := c.BodyParser(reservation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	result := database.DB.Db.Create(&reservation)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(reservation)
}

func DeleteReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation models.Reservation
	result := database.DB.Db.Where("id = ?", id).Delete(&reservation)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Reservation successfully deleted",
	})
}
