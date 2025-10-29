package presenters

import (
	"github.com/gofiber/fiber/v2"
)

type UserLogin struct {
  Token string `json:"token" bson:"token"`
  User *User  `json:"user" bson:"user"`
}

func UserLoginSuccessResponse(user *UserLogin) *fiber.Map {
	return &fiber.Map {
		"messages": []string{
			"success",
		},
		"data" : []*UserLogin{
			user,
		},
	}
}

func UserLoginMessageResponse(message string) *fiber.Map {
	return &fiber.Map {
		"messages": []string{
      message,
		},
		"data" : []*UserLogin{ },
	}
}

func UserLoginErrorResponse(err error) *fiber.Map {
	return &fiber.Map {
		"messages": []string{
			err.Error(),
		},
		"data" : []string{},
	}
}
