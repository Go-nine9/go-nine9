package main

import (
	"github.com/Go-nine9/go-nine9/controllers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
}
