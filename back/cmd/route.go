package main

import (
	"github.com/Go-nine9/go-nine9/controllers"
	"github.com/Go-nine9/go-nine9/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Unprotected route
	api.Get("/", controllers.Home)

	auth := api.Group("/auth")
	auth.Post("/register", controllers.CreateNewUser)
	auth.Post("/login", controllers.LoginUser)

	admin := api.Group("/admin")
	
	admin.Use(middleware.AuthMiddleware())

	// USERS routes are now protected by RoleMiddleware
	admin.Get("/users", middleware.RoleMiddleware("admin"), controllers.GetAllUsers)
	admin.Get("/users/:id", middleware.RoleMiddleware("admin"), controllers.GetUserById)
	admin.Post("/users", middleware.RoleMiddleware("admin"), controllers.CreateUser)
	admin.Patch("/users/:id", middleware.RoleMiddleware("admin"), controllers.UpdateUser)
	admin.Delete("/users/:id", middleware.RoleMiddleware("admin"), controllers.DeleteUser)
}
