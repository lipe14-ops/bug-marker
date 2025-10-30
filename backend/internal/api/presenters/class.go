package presenters

import (
  "github.com/gofiber/fiber/v2"
)

type Class struct {
	ID          string `json:"id" bson:"_id"`
	Name   		  string `json:"name" bson:"name"`
	Color       string `json:"color" bson:"color"`
	Description string `json:"description" bson:"description"`
}

func ClassSuccessResponse(class *Class) *fiber.Map {
  return &fiber.Map {
    "messages": []string{
      "success",
    },
    "data" : []*Class{
      class,
    },
  }
}

func ClassesSuccessResponse(class []*Class) *fiber.Map {
  return &fiber.Map {
    "messages": []string{
      "success",
    },
    "data": class,
  }
}

func ClassMessageResponse(message string) *fiber.Map {
  return &fiber.Map {
    "messages": []string{
      message,
    },
    "data" : []string{},
	}
}

func ClassErrorResponse(err error) *fiber.Map {
	return &fiber.Map {
		"messages": []string{
			err.Error(),
		},
		"data" : []string{},
	}
}

