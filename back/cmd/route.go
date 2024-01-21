package main

import (
	"github.com/Go-nine9/go-nine9/controllers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", controllers.Home)

	admin := api.Group("/admin")
	auth := api.Group("/auth")

	auth.Post("/register", controllers.CreateNewUser)
	auth.Post("/login", controllers.LoginUser)

	//USERS
	admin.Get("/users", controllers.GetAllUsers)
	admin.Get("/users/:id", controllers.GetUserById)
	admin.Post("/users", controllers.CreateUser)
	admin.Patch("/users/:id", controllers.UpdateUser)
	admin.Delete("/users/:id", controllers.DeleteUser)
}
