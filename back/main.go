package main

import (
	"github.com/go-nine9/go-nine9/db"
	"github.com/go-nine9/go-nine9/middleware"
	"github.com/go-nine9/go-nine9/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db.Connect()

	app := fiber.New()

	app.Use(middleware.AuthMiddleware())

	routes.SetupAuthRoutes(app)

	app.Listen(":8000")
}
