package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
	"backend/internal/pkg/image"
)

func CreateImage(service image.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		image, err := c.FormFile("document")

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(err))
		}

		projectID := c.Params("projectID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(errors.New("invalid project ID.")))
		}
		
		result, err := service.CreateImage(projectID, image)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(err))
		}
		
		return c.JSON(presenters.ImageSuccessResponse(result))
	}
}

func ReadImage(service image.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		imageID := c.Params("imageID")

		if imageID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(errors.New("invalid ID.")))
		}

		image, err := service.ReadImage(imageID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(err))
		}

		return c.JSON(presenters.ImageSuccessResponse(image))
	}
}

func ReadImagesByProject(service image.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projectID := c.Params("projectID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(errors.New("invalid ID.")))
		}

		images, err := service.ReadImagesByProject(projectID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(err))
		}

		return c.JSON(presenters.ImagesSuccessResponse(images))
	}
}

func UpdateImage(service image.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		imageID := c.Params("imageID")

		if imageID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(errors.New("invalid ID.")))
		}

		var requestBody entities.ImagePutRequest
		err := c.BodyParser(&requestBody)

		if requestBody.Filename == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(errors.New("invalid fields")))
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(errors.New("invalid ID.")))
		}

		image := entities.Image {
			Filename: requestBody.Filename,
		}

		result, err := service.UpdateImage(imageID, &image)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(err))
		}

		return c.JSON(presenters.ImageSuccessResponse(result))
	}
}

func DeleteImage(service image.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		imageID := c.Params("imageID")

		if imageID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ImageErrorResponse(errors.New("invalid ID.")))
		}

		err := service.DeleteImage(imageID)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.ImageErrorResponse(errors.New("image was not deleted.")))
		}

		return c.JSON(presenters.ImageMessageResponse("image was deleted."))
	}
}
