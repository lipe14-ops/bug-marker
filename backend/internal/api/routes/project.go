package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/internal/pkg/project"
	"backend/internal/api/handlers"
)

func ProjectRouter(app fiber.Router, service project.Service) {
	app.Post("/project", handlers.CreateProject(service))
	app.Get("/project/:ID", handlers.ReadProject(service))
	app.Get("/projects/owner/:ID", handlers.ReadProjectByOwner(service))
	app.Put("/project/:ID", handlers.UpdateProject(service))
	app.Delete("/project/:ID", handlers.DeleteProject(service))
}
