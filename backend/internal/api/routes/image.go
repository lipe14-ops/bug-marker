package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/internal/pkg/image"
	"backend/internal/api/handlers"
)

func ImageRouter(app fiber.Router, service image.Service) {
	app.Post("/project/:projectID/image", handlers.CreateImage(service))
	app.Get("/project/image/:imageID", handlers.ReadImage(service))
	app.Get("/project/:projectID/images", handlers.ReadImagesByProject(service))
	app.Put("/project/image/:imageID", handlers.UpdateImage(service))
	app.Delete("/project/image/:imageID", handlers.DeleteImage(service))
}
