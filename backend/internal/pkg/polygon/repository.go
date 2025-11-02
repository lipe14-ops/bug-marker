package polygon

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
)

type Repository interface {
	CreatePolygon(imageID, classID string, polygon *entities.Polygon) (*presenters.Polygon, error)
	ReadPolygon(imageID, polygonID string) (*presenters.Polygon, error)
  ReadPolygonsByImage(ID string) ([]*presenters.Polygon, error)
  UpdatePolygon(imageID, polygonID, classID string, polygon *entities.Polygon) (*presenters.Polygon, error)
  DeletePolygon(imageID, polygonID string) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository {
		collection: collection,
	}
}

func (r *repository) CreatePolygon(imageID, classID string, polygon *entities.Polygon) (*presenters.Polygon, error) {
  imageObjectID, err := primitive.ObjectIDFromHex(imageID)

	if err != nil {
		return nil, err
	}

  classObjectID, err := primitive.ObjectIDFromHex(classID)

	if err != nil {
		return nil, err
	}

  polygon.ID = primitive.NewObjectID()
  polygon.ImageID = imageObjectID 
  polygon.ClassID = classObjectID 
  polygon.Type = "Polygon"
  polygon.CreatedAt = time.Now()
  
  _, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": imageObjectID }, bson.M { "$push": bson.M { "polygons": polygon } })
	
	if err != nil {
		return nil, err
	}

	responsePolygon := &presenters.Polygon {
    ID: polygon.ID.Hex(),
    Name: polygon.Name,
    ImageID: polygon.ImageID.Hex(),
    ClassID: polygon.ClassID.Hex(),
    Type: polygon.Type,
    Coordinates: polygon.Coordinates,
	}

	return responsePolygon, err
}

func (r *repository) ReadPolygonsByImage(ID string) ([]*presenters.Polygon, error) {
	imageID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	filter := bson.M {
		"_id": imageID,
	}

	var image entities.Image
	err = r.collection.FindOne(context.Background(), filter).Decode(&image)

	if err != nil {
		return nil, errors.New("image was not found.")
	}

  var responsePolygons []*presenters.Polygon

  for _, polygon := range image.Polygons {
    responsePolygon := &presenters.Polygon {
      ID: polygon.ID.Hex(),
      Name: polygon.Name,
      ImageID: polygon.ImageID.Hex(),
      ClassID: polygon.ClassID.Hex(),
      Type: polygon.Type,
      Coordinates: polygon.Coordinates,
    }

    responsePolygons = append(responsePolygons, responsePolygon)
  }

  return responsePolygons, nil
}

func (r *repository) ReadPolygon(imageID, polygonID string) (*presenters.Polygon, error) {
	imageObjectID, err := primitive.ObjectIDFromHex(imageID)

	if err != nil {
		return nil, err
	}

	filter := bson.M {
		"_id": imageObjectID,
	}

	var image entities.Image
	err = r.collection.FindOne(context.Background(), filter).Decode(&image)

	if err != nil {
		return nil, errors.New("image was not found.")
	}

  for _, polygon := range image.Polygons {
    if polygon.ID.Hex() == polygonID {
      responsePolygon := &presenters.Polygon {
        ID: polygon.ID.Hex(),
        Name: polygon.Name,
        ImageID: polygon.ImageID.Hex(),
        ClassID: polygon.ClassID.Hex(),
        Type: polygon.Type,
        Coordinates: polygon.Coordinates,
      }
      return responsePolygon, nil
    }
  }

  return nil, errors.New("polygon was not found.")
}

func (r *repository) UpdatePolygon(imageID, polygonID, classID string, polygon *entities.Polygon) (*presenters.Polygon, error) {
	imageObjectID, err := primitive.ObjectIDFromHex(imageID)

	if err != nil {
		return nil, err
	}

	polygonObjectID, err := primitive.ObjectIDFromHex(polygonID)

	if err != nil {
		return nil, err
	}

	classObjectID, err := primitive.ObjectIDFromHex(classID)

	if err != nil {
		return nil, err
	}

  polygon.ImageID = imageObjectID
  polygon.ClassID = classObjectID
  polygon.ID = polygonObjectID
  polygon.Type = "Polygon"

  _, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": imageObjectID, "polygons._id": polygonObjectID }, bson.M { "$set": bson.M { "polygons.$": polygon }})
	
	if err != nil {
		return nil, err
	}

  responsePolygon := &presenters.Polygon {
    ID: polygonID,
    Name: polygon.Name,
    ImageID: imageID,
    ClassID: polygon.ClassID.Hex(),
    Type: polygon.Type,
    Coordinates: polygon.Coordinates,
  }

  return responsePolygon, nil
}

func (r *repository) DeletePolygon(imageID, polygonID string) error {
	imageObjectID, err := primitive.ObjectIDFromHex(imageID)

	if err != nil {
		return err
	}

	polygonObjectID, err := primitive.ObjectIDFromHex(polygonID)

	if err != nil {
		return err
	}

  _, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": imageObjectID }, bson.M { "$pull": bson.M { "polygons": bson.M { "_id": polygonObjectID }}})

	if err != nil {
		return err
	}

	return nil
}
