package routes

import (
	"github.com/go-nine9/go-nine9/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/api/register/user", fiber.Handler(controller.RegisterUser))
	app.Post("/api/login", fiber.Handler(controller.Login))
}
