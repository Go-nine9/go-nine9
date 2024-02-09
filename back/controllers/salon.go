package controllers

import (
	"net/http"

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/helper"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

type SalonRequest struct {
	User    []models.User `json:"user"`
	Phone   string        `json:"phone"`
	Name    string        `json:"name"`
	Address string        `json:"address"`
}

func CreateSalon(c *fiber.Ctx) error {
	// Vérifie si les revendications de l'utilisateur existent
	claims, err := c.Locals("userClaims").(jwt.MapClaims)
	if !err {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	// Parse la requete
	var salonRequest SalonRequest
	if err := c.BodyParser(&salonRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	salon := models.Salon{
		Name:    salonRequest.Name,
		Address: salonRequest.Address,
		Phone:   salonRequest.Phone,
	}

	//génère un UUID pour le salon
	salon.ID, _ = models.GenerateUUID()

	//crée le salon
	result := database.DB.Db.Create(&salon)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	// récupère l'id de la personne qui créé le salon
	userID := claims["id"].(string)

	// Convert the manager ID from string to UUID
	managerID, _ := uuid.Parse(userID)

	// Find the manager by ID
	var manager models.User
	result = database.DB.Db.First(&manager, "id = ?", managerID)

	// Check if there is an error and the manager doesn't exist
	if result.Error != nil && result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Manager not found",
		})
	}

	role := claims["role"].(string)
	// Associate the salon to manager or create a new manager if not found
	manager.SalonID = &salon.ID

	// Vérifie si l'utilisateur est déjà un employé d'un salon
	if role == "employee" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "You already have a salon",
		})
	}

	// Vérifie si l'utilisateur a un rôle utilisateur
	if role == "users" {
		manager.Roles = "manager"

	}

	// Met à jour l'utilisateur dans la base de données avec le nouveau rôle et l'ID du salon
	if err := database.DB.Db.Save(&manager).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messageHooo": err.Error(),
		})
	}

	signedToken, ok := helper.GenerateToken(manager.ID, manager.Roles, manager.Firstname, manager.Email, salon.ID)

	if ok != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not ",
		})
	}
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

		if err := database.DB.Db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"messageHooo": err.Error(),
			})
		}

	}

	response := fiber.Map{
		"jwt": signedToken,
	}

	return c.Status(fiber.StatusOK).JSON(response)

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
