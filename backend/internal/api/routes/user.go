package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/internal/pkg/user"
	"backend/internal/api/handlers"
)

func UserRouter(app fiber.Router, service user.Service) {
	app.Post("/user", handlers.CreateUser(service))
	app.Get("/user/:ID", handlers.ReadUser(service))
	app.Put("/user/:ID", handlers.UpdateUser(service))
	app.Delete("/user/:ID", handlers.DeleteUser(service))
}
