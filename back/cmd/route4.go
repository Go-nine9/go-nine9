package main

// import (
// 	"github.com/Go-nine9/go-nine9/controllers"
// 	"github.com/Go-nine9/go-nine9/middleware"
// 	"github.com/gofiber/fiber/v2"
// )

// func setupRoutes(app *fiber.App) {

// 	// NEW VERSION
// 	// Public routes
// 	app.Get("/", controllers.Home)

// 	app.Get("/register", controllers.CreateNewUser)
// 	app.Get("/login", controllers.LoginUser)
// 	app.Post("/register", controllers.CreateNewUser)
// 	app.Post("/login", controllers.LoginUser)

// 	app.Get("/salons", controllers.GetSalons)
// 	app.Get("/salons/:id", controllers.GetSalonById)

// 	// Protected routes
// 	api := app.Group("/api")
// 	api.Use(middleware.AuthMiddleware())

// 	// ADMIN ONLY
// 	admin := api.Group("/admin")
// 	admin.Use(middleware.RoleMiddleware("admin"))

// 	// USERS
// 	admin.Get("/users", controllers.GetAllUsers)
// 	admin.Get("/users/:id", controllers.GetUserById)
// 	admin.Post("/users", controllers.CreateUser)
// 	admin.Patch("/users/:id", controllers.UpdateUser)
// 	admin.Delete("/users/:id", controllers.DeleteUser)

// 	//SALONS
// 	admin.Post("/salons", controllers.CreateSalon)
// 	admin.Post("/salons/:id", controllers.AddStaff)
// 	admin.Put("/salons/:id", controllers.UpdateSalon)
// 	admin.Delete("/salons/:id", controllers.DeleteSalon)

// 	// SLOTS

// 	// MANAGER ONLY
// 	manager := api.Group("/manager")
// 	manager.Use(middleware.RoleMiddleware("manager"))

// // OLD VERSION //
// api := app.Group("/api")
// // Unprotected route
// api.Get("/", controllers.Home)

// auth.Post("/register", controllers.CreateNewUser)
// auth.Post("/login", controllers.LoginUser)

// auth.Post("/register", controllers.CreateNewUser)
// auth.Post("/login", controllers.LoginUser)

// /// AUTH ///
// api.Get("/reservations", controllers.GetAllReservation)
// api.Get("/reservations/:id", controllers.GetReservationById)

// /// ADMIN ONLY ///
// /// AUTH ///
// api.Get("/reservations", controllers.GetAllReservation)
// api.Get("/reservations/:id", controllers.GetReservationById)

// /// ADMIN ONLY ///
// admin := api.Group("/admin")

// admin.Use(middleware.AuthMiddleware())

// // USERS routes are now protected by RoleMiddleware
// admin.Get("/users", middleware.RoleMiddleware("admin"), controllers.GetAllUsers)
// admin.Get("/users/:id", middleware.RoleMiddleware("admin"), controllers.GetUserById)
// admin.Post("/users", middleware.RoleMiddleware("admin"), controllers.CreateUser)
// admin.Patch("/users/:id", middleware.RoleMiddleware("admin"), controllers.UpdateUser)
// admin.Delete("/users/:id", middleware.RoleMiddleware("admin"), controllers.DeleteUser)

// //SALONS
// api.Post("/salons", controllers.CreateSalon)
// api.Get("/salons", controllers.GetSalons)
// api.Get("/salons/:id", controllers.GetSalonById)
// api.Put("/salons/:id", controllers.UpdateSalon)
// api.Delete("/salons/:id", controllers.DeleteSalon)
// // Add new staff member to the salon
// api.Post("/salons/:id", controllers.AddStaff)

// /// MANAGER ONLY ///
// manager := api.Group("/manager")
// //SLOTS
// manager.Get("/slots", controllers.GetAllOwnSlots)
// }
