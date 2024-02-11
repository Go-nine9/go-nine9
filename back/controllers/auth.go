package controllers

import (
	"os"
	"time"

	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = os.Getenv("JWT_SECRET")

func CreateNewUser(c *fiber.Ctx) error {
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
	user.Password = hashedPassword

	user.ID, err = models.GenerateUUID()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate UUID",
		})
	}

	if user.Roles == "" {
		user.Roles = "users"
	}

	database.DB.Db.Create(&user)

	//IF user is a manager, create a salon
	if user.Roles == "manager" {

		salon := new(models.Salon)
		salon.ID, err = models.GenerateUUID()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not generate UUID",
			})
		}
		salon.Name = user.Firstname + "'s salon"
		database.DB.Db.Create(&salon)
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 1 jour
		"role":      user.Roles,
		"firstname": user.Firstname,
		"email":     user.Email,
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}
	userWithToken := fiber.Map{
		"jwt": token,
	}
	return c.Status(200).JSON(userWithToken)
}

func LoginUser(c *fiber.Ctx) error {

	loginUser := new(models.User)

	if err := c.BodyParser(loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var existingUser models.User
	if err := database.DB.Db.Where("email = ?", loginUser.Email).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := models.VerifyPassword(existingUser.Password, loginUser.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        existingUser.ID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 1 jour
		"role":      existingUser.Roles,
		"firstname": existingUser.Firstname,
		"email":     existingUser.Email,
		"salonID":   existingUser.SalonID,
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not generate token",
		})
	}
	response := fiber.Map{
		"jwt": token,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
