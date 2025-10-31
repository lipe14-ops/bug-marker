package presenters

import (
	"github.com/gofiber/fiber/v2"
)

type Image struct {
	ID         string `json:"id" bson:"_id"`
	ProjectID  string `json:"projectID" bson:"projectID"`
	Filename	 string `json:"filename" bson:"filename"`
	Url        string `json:"url" bson:"url"`
}

func ImageSuccessResponse(image *Image) *fiber.Map {
	return&fiber.Map {
		"messages": []string{
			"success",
		},
		"data" : []*Image{
			image,
		},
	}
}

func ImagesSuccessResponse(images []*Image) *fiber.Map {
	return&fiber.Map {
		"messages": []string{
			"success",
		},
		"data": images,
	}
}

func ImageMessageResponse(message string) *fiber.Map {
	return&fiber.Map {
		"messages": []string{
			message,
		},
		"data" : []string{},
	}
}

func ImageErrorResponse(err error) *fiber.Map {
	return &fiber.Map {
		"messages": []string{
			err.Error(),
		},
		"data" : []string{},
	}
}
