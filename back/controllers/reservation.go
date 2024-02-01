package controllers

import (
	"fmt"

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
)

func GetReservationById(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation models.Reservation

	result := database.DB.Db.Where("id = ?", id).First(&reservation)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return c.JSON(reservation)
}

func GetAllReservation(c *fiber.Ctx) error {
	// id := "9f7ceeaf-7132-4cee-a660-ce801349853c"
	// id_slot := "9f7ceeaf-7132-4cee-a660-ce801349854c"
	var reservation []models.Reservation

	// result := database.DB.Db.Model(&reservation).Preload("User", func(db *gorm.DB) *gorm.DB {
	// 	return db.Select("ID, Lastname, Firstname")
	// }).Where("customer_id = ?", id).First(&reservation)

	// result := database.DB.Db.Model(&reservation).
	// 	Joins("LEFT JOIN slots ON reservations.slot_id = slots.id").
	// 	Joins("LEFT JOIN users ON reservations.customer_id = users.id").
	// 	Select("users.id, users.email, reservations.id, slots.id, slots.date").
	// 	Where("reservations.customer_id = ?", id).
	// 	Find(&reservation)
	// result := database.DB.Db.Model(&reservation).Preload("User.Salon").Preload("Slot").Where("slot_id = ?", id_slot).Find(&reservation)
	// result := database.DB.Db.Model(&reservation).Preload("User").Preload("Slot").Preload("Slot.HairdressingStaff").Where("slot_id = ?", id_slot).Find(&reservation)
	result := database.DB.Db.
		Preload("User").
		Preload("Slot").
		Preload("Slot.HairdressingStaff").
		Preload("Slot.HairdressingStaff.Salon").
		Where("reservations.slot_id = ?", "9f7ceeaf-7132-4cee-a660-ce801349854c").
		Find(&reservation)

	if result.Error != nil {
		// Gérer l'erreur
		fmt.Println(result.Error)
	} else {
		// Les informations demandées sont dans la variable reservations
		for _, reservation := range reservation {
			fmt.Println("ID de la réservation:", reservation.ID)
			fmt.Println("ID du client (Customer):", reservation.User.ID)
			fmt.Println("Nom du client (Customer):", reservation.User.Lastname)
			fmt.Println("Prénom du client (Customer):", reservation.User.Firstname)
			fmt.Println("ID du coiffeur (HairdressingStaffID):", reservation.Slot.HairdressingStaffID)
			fmt.Println("Nom du coiffeur (HairdressingStaff):", reservation.Slot.HairdressingStaff.Lastname)
			fmt.Println("NOM du salon (HairdressingStaff.Salon):", reservation.Slot.HairdressingStaff.Salon.Name)
			fmt.Println("Date du créneau (Slot):", reservation.Slot.Date)
			fmt.Println("Heure de début du créneau (Slot):", reservation.Slot.BeginningHour)
			fmt.Println("Heure de fin du créneau (Slot):", reservation.Slot.EndTime)
		}
	}

	return c.JSON(reservation)
}
