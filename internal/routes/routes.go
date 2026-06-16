package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/go-backend-task/internal/handler"
	"github.com/yourusername/go-backend-task/internal/middleware"
)

// SetupRoutes registers application routes and middleware.
func SetupRoutes(app *fiber.App, handler *handler.UserHandler) {
	api := app.Group("/api/v1")
	api.Use(middleware.RequestID)
	api.Use(middleware.RequestLogger)

	api.Post("/users", handler.CreateUser)
	api.Get("/users", handler.ListUsers)
	api.Get("/users/:id", handler.GetUserByID)
	api.Put("/users/:id", handler.UpdateUser)
	api.Delete("/users/:id", handler.DeleteUser)
}
