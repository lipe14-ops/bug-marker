package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/internal/pkg/auth"
	"backend/internal/api/handlers"
)

func AuthRoutes(app fiber.Router, service auth.Service) {
	app.Post("/login", handlers.LoginHandler(service))
}
