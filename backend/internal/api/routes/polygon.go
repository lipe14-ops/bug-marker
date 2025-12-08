package routes

import (
	"backend/internal/api/handlers"
	"backend/internal/pkg/polygon"

	"github.com/gofiber/fiber/v2"
)

func PolygonRouter(app fiber.Router, service polygon.Service) {
	app.Post("/project/image/:imageID/polygon", handlers.CreatePolygon(service))
	app.Get("/project/image/:imageID/polygon/:polygonID", handlers.ReadPolygon(service))
	app.Get("/project/image/:ID/polygons", handlers.ReadPolygonsByImage(service))
	app.Put("/project/image/:imageID/polygon/:polygonID", handlers.UpdatePolygon(service))
	app.Delete("/project/image/:imageID/polygon/:polygonID", handlers.DeletePolygon(service))
}
