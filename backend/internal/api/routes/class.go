package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/internal/pkg/class"
	"backend/internal/api/handlers"
)

func ClassRouter(app fiber.Router, service class.Service) {
  app.Post("/project/:projectID/class", handlers.CreateClass(service))
	app.Get("/project/:projectID/class/:classID", handlers.ReadClass(service))
	app.Get("/project/:ID/classes", handlers.ReadClassByProject(service))
  app.Put("/project/:ID/class", handlers.UpdateClass(service))
	app.Delete("/project/:projectID/class/:classID", handlers.DeleteClass(service))
}
