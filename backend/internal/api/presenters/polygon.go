package presenters

import (
	"github.com/gofiber/fiber/v2"
)

type Polygon struct {
	ID              string `json:"id" bson:"_id"`
	Name            string `json:"name" bson:"name"`
	ImageID         string `json:"imageID" bson:"imageID"`
	ClassID         string `json:"classID" bson:"classID"`
	Type            string `json:"type" bson:"type"`
	Coordinates     [][][]float64 `json:"coordinates" bson:"coordinates"`
}

func polygonAdapter(polygon *Polygon) *fiber.Map {
  return &fiber.Map {
    "id": polygon.ID,
    "name": polygon.Name,
    "imageID": polygon.ImageID,
    "classID": polygon.ClassID,
    "type": polygon.Type,
    "coordinates": polygon.Coordinates[0],
    }
}

func PolygonSuccessResponse(polygon *Polygon) *fiber.Map {
  return &fiber.Map {
    "messages": []string{
      "success",
    },
    "data" : []*fiber.Map {
      polygonAdapter(polygon),
    },
  }
}

func PolygonsSuccessResponse(polygons []*Polygon) *fiber.Map {
  polys := []*fiber.Map{}

  for _, poly := range polygons {
    polys = append(polys, polygonAdapter(poly))
  }

  return &fiber.Map {
    "messages": []string{
      "success",
    },
    "data": polys,
  }
}

func PolygonMessageResponse(message string) *fiber.Map {
  return &fiber.Map {
    "messages": []string{
      message,
    },
    "data" : []string{},
	}
}

func PolygonErrorResponse(err error) *fiber.Map {
	return &fiber.Map {
		"messages": []string{
			err.Error(),
		},
		"data" : []string{},
	}
}

