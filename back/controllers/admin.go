package controllers

import (
	"github.com/Go-nine9/go-nine9/database"
	"github.com/Go-nine9/go-nine9/models" // import the models package
	"github.com/gofiber/fiber/v2"
)

func AdminDashboard(c *fiber.Ctx) error {
	return c.SendString("Welcome to the admin dashboard!")
}

// / USERS ///
// func GetAllUsers(c *fiber.Ctx) error {
// 	var users []models.User
// 	database.DB.Db.Find(&users)
// 	return c.JSON(users)
// }

// func GetUserById(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var user models.User
// 	database.DB.Db.Where("id = ?", id).First(&user)
// 	return c.JSON(user)
// }

// func CreateUser(c *fiber.Ctx) error {
// 	user := new(models.User)
// 	if err := c.BodyParser(user); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	hashedPassword, err := models.HashPassword(user.Password)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to hash password",
// 		})
// 	}
// 	user.ID, _ = models.GenerateUUID()

// 	user.Password = hashedPassword

// 	database.DB.Db.Create(&user)
// 	return c.Status(200).JSON(user)
// }

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
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

	result := database.DB.Db.Where("id = ?", id).Updates(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return c.SendString("User successfully updated")
}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}
	return c.SendString("User successfully deleted")
}
