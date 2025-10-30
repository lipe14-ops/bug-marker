package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/internal/api/presenters"
	"backend/internal/pkg/class"
	"backend/internal/pkg/entities"
)

func CreateClass(service class.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.ClassPostRequest
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(err))
		}

		if requestBody.Name == "" || requestBody.Color == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid fields")))
		}

		projectID := c.Params("projectID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid project id")))
		}

		class := entities.Class {
			Name: requestBody.Name,
			Description: requestBody.Description,
			Color: requestBody.Color,
		}

		result, err := service.CreateClass(projectID, &class)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(err))
		}
		
		return c.JSON(presenters.ClassSuccessResponse(result))
	}
}

func ReadClass(service class.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		classID := c.Params("classID")

		if classID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid class ID.")))
		}

		projectID := c.Params("projectID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid project ID.")))
		}

		user, err := service.ReadClass(projectID, classID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(err))
		}

		return c.JSON(presenters.ClassSuccessResponse(user))
	}
}

func ReadClassByProject(service class.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projectID := c.Params("ID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid ID.")))
		}

		classes, err := service.ReadClassesByProject(projectID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(err))
		}

		return c.JSON(presenters.ClassesSuccessResponse(classes))
	}
}

func UpdateClass(service class.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projectID := c.Params("ID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid class ID.")))
		}

		var requestBody entities.ClassPutRequest
		err := c.BodyParser(&requestBody)

		if requestBody.Name == "" || requestBody.Description == "" || requestBody.Color == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid fields")))
		}
    
		class := entities.Class {
      ID: requestBody.ID,
			Name: requestBody.Name,
      Description: requestBody.Description,
      Color: requestBody.Color,
		}

		result, err := service.UpdateClass(projectID, &class)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(err))
		}

		return c.JSON(presenters.ClassSuccessResponse(result))
	}
}

func DeleteClass(service class.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		classID := c.Params("classID")

		if classID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid class ID.")))
		}

		projectID := c.Params("projectID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ClassErrorResponse(errors.New("invalid project ID.")))
		}

		err := service.DeleteClass(projectID, classID)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.ClassErrorResponse(errors.New("class was not deleted.")))
		}

		return c.JSON(presenters.ClassMessageResponse("class was deleted."))
	}
}
