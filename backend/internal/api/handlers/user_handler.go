package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
	"backend/internal/pkg/user"
)

func CreateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.UserPostRequest
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		if requestBody.Name == "" || requestBody.Email == "" || requestBody.Password == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid fields")))
		}

		user := entities.User {
			Name: requestBody.Name,
			Email: requestBody.Email,
			Password: requestBody.Password,
		}

		result, err := service.CreateUser(&user)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(err))
		}
		
		return c.JSON(presenters.UserSuccessResponse(result))
	}
}

func ReadUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("ID")

		if userID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid ID.")))
		}

		user, err := service.ReadUser(userID)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		return c.JSON(presenters.UserSuccessResponse(user))
	}
}

func UpdateUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("ID")

		if userID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid ID.")))
		}

		var requestBody entities.UserPutRequest
		err := c.BodyParser(&requestBody)

		if requestBody.Name == "" || requestBody.Email == "" || requestBody.Password == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid fields")))
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid ID.")))
		}

		user := entities.User {
			Name: requestBody.Name,
			Email: requestBody.Email,
			Password: requestBody.Password,
		}

		result, err := service.UpdateUser(userID, &user)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		return c.JSON(presenters.UserSuccessResponse(result))
	}
}

func DeleteUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Params("ID")

		if userID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(errors.New("invalid ID.")))
		}

		err := service.DeleteUser(userID)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.UserErrorResponse(errors.New("user was not deleted.")))
		}

		return c.JSON(presenters.UserMessageResponse("user was deleted."))
	}
}
