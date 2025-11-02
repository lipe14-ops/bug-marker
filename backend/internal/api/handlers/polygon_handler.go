package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/internal/api/presenters"
	"backend/internal/pkg/polygon"
	"backend/internal/pkg/entities"
)

func CreatePolygon(service polygon.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.PolygonPostRequest
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(err))
		}

		imageID := c.Params("imageID")

		if imageID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid image id")))
		}

		if requestBody.ClassID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid class id")))
		}

		polygon := entities.Polygon {
      Name: requestBody.Name,
      Coordinates: [][][]float64{ requestBody.Coordinates },
		}

		result, err := service.CreatePolygon(imageID, requestBody.ClassID, &polygon)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.PolygonErrorResponse(err))
		}
		
		return c.JSON(presenters.PolygonSuccessResponse(result))
	}
}

func ReadPolygon(service polygon.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		imageID := c.Params("imageID")

		if imageID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid image ID.")))
		}

		polygonID := c.Params("polygonID")

		if polygonID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid class ID.")))
		}

		user, err := service.ReadPolygon(imageID, polygonID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.PolygonErrorResponse(err))
		}

		return c.JSON(presenters.PolygonSuccessResponse(user))
	}
}

func ReadPolygonsByImage(service polygon.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		imageID := c.Params("ID")

		if imageID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid ID.")))
		}

		polygons, err := service.ReadPolygonsByImage(imageID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.PolygonErrorResponse(err))
		}

		return c.JSON(presenters.PolygonsSuccessResponse(polygons))
	}
}

func UpdatePolygon(service polygon.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		imageID := c.Params("imageID")

		if imageID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.PolygonErrorResponse(errors.New("invalid image ID.")))
		}

		polygonID := c.Params("polygonID")

		if polygonID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.PolygonErrorResponse(errors.New("invalid polygon ID.")))
		}

		var requestBody entities.PolygonPutRequest
		err := c.BodyParser(&requestBody)

		if requestBody.ClassID == "" || requestBody.Coordinates == nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.PolygonErrorResponse(errors.New("invalid fields")))
		}
    
		polygon := entities.Polygon {
      Name: requestBody.Name,
      Coordinates: [][][]float64{ requestBody.Coordinates },
		}

		result, err := service.UpdatePolygon(imageID, polygonID, requestBody.ClassID, &polygon)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.PolygonErrorResponse(err))
		}

		return c.JSON(presenters.PolygonSuccessResponse(result))
	}
}

func DeletePolygon(service polygon.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		imageID := c.Params("imageID")

		if imageID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid class ID.")))
		}

		polygonID := c.Params("polygonID")

		if polygonID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid project ID.")))
		}

		err := service.DeletePolygon(imageID, polygonID)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.ClassErrorResponse(errors.New("polygon was not deleted.")))
		}

		return c.JSON(presenters.ClassMessageResponse("polygon was deleted."))
	}
}
