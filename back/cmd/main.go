package main

import (
	"github.com/Go-nine9/go-nine9/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	setupRoutes(app)
	app.Listen(":8097")
}
