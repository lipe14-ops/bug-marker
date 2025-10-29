package handlers

import (
  "time"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
	"backend/internal/pkg/auth"

  "github.com/golang-jwt/jwt/v5"
)


func LoginHandler(service auth.Service) fiber.Handler {
  return func (c *fiber.Ctx) error {
    var userLogin entities.LoginPost
		err := c.BodyParser(&userLogin)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.UserLoginErrorResponse(err))
		}

    user, err := service.ReadUserByCredentials(userLogin.Email, userLogin.Password) 

    if err != nil  {
		 c.Status(http.StatusBadRequest)
		 return c.JSON(presenters.UserLoginMessageResponse("invalid credentials."))
    }

    claims := jwt.MapClaims {
     "id":        user.ID,
     "exp":   time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    t, err := token.SignedString([]byte("secret"))

    if err != nil {
     return c.SendStatus(http.StatusInternalServerError)
    }

    userLoginResponse := &presenters.UserLogin {
      Token: t,
      User: user,
    }

    return c.JSON(presenters.UserLoginSuccessResponse(userLoginResponse))
  }
}
