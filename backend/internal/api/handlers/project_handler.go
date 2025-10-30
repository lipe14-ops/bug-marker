package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
	"backend/internal/pkg/project"
)

func CreateProject(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.ProjectPostRequest
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(err))
		}

		if requestBody.Owner == "" || requestBody.Title == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid fields")))
		}
    
		project := entities.Project {
      Title: requestBody.Title,
      Description: requestBody.Description,
		}

		result, err := service.CreateProject(requestBody.Owner, &project)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(err))
		}
		
		return c.JSON(presenters.ProjectSuccessResponse(result))
	}
}

func ReadProject(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projectID := c.Params("ID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid ID.")))
		}

		project, err := service.ReadProject(projectID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(err))
		}

		return c.JSON(presenters.ProjectSuccessResponse(project))
	}
}

func ReadProjectByOwner(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ownerID := c.Params("ID")

		if ownerID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid ID.")))
		}

		projects, err := service.ReadProjectByOwner(ownerID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(err))
		}

		return c.JSON(presenters.ProjectsSuccessResponse(projects))
	}
}

func UpdateProject(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projectID := c.Params("ID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid ID.")))
		}

		var requestBody entities.ProjectPutRequest
		err := c.BodyParser(&requestBody)

		if requestBody.Title == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid fields")))
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid ID.")))
		}

		project := entities.Project {
			Title: requestBody.Title,
			Description: requestBody.Description,
		}

		result, err := service.UpdateProject(projectID, &project)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(err))
		}

		return c.JSON(presenters.ProjectSuccessResponse(result))
	}
}

func DeleteProject(service project.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		projectID := c.Params("ID")

		if projectID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("invalid ID.")))
		}

		err := service.DeleteProject(projectID)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.ProjectErrorResponse(errors.New("project was not deleted.")))
		}

		return c.JSON(presenters.ProjectMessageResponse("project was deleted."))
	}
}
