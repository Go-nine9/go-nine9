package main

import (
	"github.com/Go-nine9/go-nine9/controllers"
	"github.com/Go-nine9/go-nine9/middleware"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	// PUBLIC ROUTES
	app.Get("/", controllers.Home)
	app.Get("/salons", controllers.GetSalons)
	app.Get("/salons/:id", controllers.GetSalonById)

	// AUTH ROUTES
	app.Post("/auth/register", controllers.CreateNewUser)
	app.Post("/auth/login", controllers.LoginUser)

	// PRIVATE ROUTES
	api := app.Group("/api", middleware.AuthRequired())
	api.Get("/me", controllers.GetMe)
	api.Patch("/me", controllers.UpdateMe)
	api.Patch("/me/password", controllers.UpdateMePassword)

	// GESTION ROUTES (ADMIN AND MANAGER)
	management := api.Group("/management", middleware.RoleMiddleware("manager", "admin"))
	management.Post("/users", controllers.CreateUser)

	//EMPLOYEE ONLY ROUTES
	// employee := api.Group("/employee")
	// employee.Use(controllers.RoleMiddleware("employee"))

	// //SLOTS ACTIONS
	// employee.Get("/salons/slots", controllers.GetEmployeeSlots)
	// // MANAGER ONLY ROUTES
	// manager := api.Group("/manager")
	// manager.Use(controllers.RoleMiddleware("manager"))

	// //SALONS ACTIONS
	// manager.Get("/salons", controllers.GetMySalons)
	// manager.Post("/salons", controllers.CreateMySalon)
	// manager.Patch("/salons/:id", controllers.UpdateMySalon)
	// manager.Delete("/salons/:id", controllers.DeleteMySalon)

	// //STAFF ACTIONS
	// manager.Get("/salons/staff", controllers.GetManagerStaff)
	// manager.Post("/salons/staff", controllers.AddManagerStaff)
	// manager.Patch("/salons/staff/:id", controllers.UpdateManagerStaff)
	// manager.Delete("/salons/staff/:id", controllers.DeleteManagerStaff)

	// //SLOTS ACTIONS
	// manager.Get("/salons/slots", controllers.GetManagerSlots)
	// manager.Post("/salons/slots", controllers.CreateManagerSlot)
	// manager.Patch("/salons/slots/:id", controllers.UpdateManagerSlot)
	// manager.Delete("/salons/slots/:id", controllers.DeleteManagerSlot)

	// //RESERVATIONS ACTIONS
	// manager.Get("/salons/reservations", controllers.GetManagerReservations)
	// manager.Post("/salons/reservations", controllers.CreateManagerReservation)
	// manager.Patch("/salons/reservations/:id", controllers.UpdateManagerReservation)
	// manager.Delete("/salons/reservations/:id", controllers.DeleteManagerReservation)

	// // ADMIN ONLY ROUTES
	// admin := api.Group("/admin")
	// admin.Use(controllers.RoleMiddleware("admin"))

	// //USERS ACTIONS
	// admin.Get("/users", controllers.GetAllUsers)
	// admin.Get("/users/:id", controllers.GetUserById)
	// admin.Post("/users", controllers.CreateUser)
	// admin.Patch("/users/:id", controllers.UpdateUser)
	// admin.Delete("/users/:id", controllers.DeleteUser)

	// //SALONS ACTIONS
	// admin.Post("/salons", controllers.CreateSalon)
	// admin.Patch("/salons/:id", controllers.UpdateSalon)
	// admin.Delete("/salons/:id", controllers.DeleteSalon)
	// admin.Post("/salons/:id", controllers.AddStaff)

	// //SLOTS ACTIONS
	// admin.Post("/salon/:id/slots", controllers.CreateSlot)
	// admin.Patch("/salon/slots/:slot_id", controllers.UpdateSlot)
	// admin.Delete("/salon/slots/:slot_id", controllers.DeleteSlot)

	// //RESERVATIONS ACTIONS
	// admin.Get("/salon/:id/reservations", controllers.GetAllReservations)
	// admin.Post("/salon/:id/reservations", controllers.CreateReservation)
	// admin.Patch("/salon/reservations/:reservation_id", controllers.UpdateReservation)
	// admin.Delete("/salon/reservations/:reservation_id", controllers.DeleteReservation)

}
