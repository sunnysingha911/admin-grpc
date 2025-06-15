package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sunnysingha911/admin-service/handlers"
)

func RegisterUserRoutes(api fiber.Router, handler *handlers.UserHandler) {
	// Group routes under /api/v1/users
	users := api.Group("/users")

	// Public route for login
	users.Post("/login", handler.Login)

	// Admin routes for user management
	users.Get("/", handler.ListUsers) // GET /api/v1/users
	users.Get("/all-users", handler.GetAllUser)
	users.Get("/:id", handler.GetUser)       // GET /api/v1/users/:id
	users.Put("/:id", handler.UpdateUser)    // PUT /api/v1/users/:id
	users.Delete("/:id", handler.DeleteUser) // DELETE /api/v1/users/:id
}
