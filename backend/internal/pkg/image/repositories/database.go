package repositories

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
	CreateImage(ID string, image *entities.Image) (*presenters.Image, error)
	ReadImage(ID string) (*presenters.Image, error)
	ReadImagesByProject(ID string) ([]*presenters.Image, error)
	UpdateImage(ID string, image *entities.Image) (*presenters.Image, error)
	DeleteImage(ID string) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository {
		collection: collection,
	}
}

func (r *repository) CreateImage(ID string, image *entities.Image) (*presenters.Image, error) {
  projectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	image.ID = primitive.NewObjectID()
	image.HasBeenSoftDeleted = false
	image.ProjectID = projectID
	image.CreatedAt = time.Now()

	_, err = r.collection.InsertOne(context.Background(), image)

	if err != nil {
		return nil, err
	}

	responseImage := &presenters.Image {
		ID: image.ID.Hex(),
		ProjectID: ID,
		Filename: image.Filename,
		Url: "foto url",
	}

	return responseImage, err
}

func (r *repository) ReadImagesByProject(ID string) ([]*presenters.Image, error) {
	projectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	filter := bson.M {
		"projectID": projectID,
	}

  cursor, err := r.collection.Find(context.TODO(), filter)

	var images []entities.Image
  if err = cursor.All(context.TODO(), &images); err != nil {
    return nil, err
  }

  var responseImages []*presenters.Image
  
  for _, image := range images {
		responseImage := &presenters.Image {
			ID: image.ID.Hex(),
			ProjectID: image.ProjectID.Hex(),
			Filename: image.Filename,
			Url: "foto url",
		}

		responseImages = append(responseImages, responseImage)
  }

	return responseImages, nil
}

func (r *repository) ReadImage(ID string) (*presenters.Image, error) {
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

	responseImage := &presenters.Image {
		ID: image.ID.Hex(),
		ProjectID: image.ProjectID.Hex(),
		Filename: image.Filename,
		Url: "foto url",
	}

	return responseImage, nil
}

func (r *repository) UpdateImage(ID string, image *entities.Image) (*presenters.Image, error) {
	imageID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	image.ID = imageID
	
	_, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": imageID }, bson.M { "$set": image })
	
	if err != nil {
		return nil, err
	}

	responseImage := &presenters.Image {
		ID: image.ID.Hex(),
		ProjectID: image.ProjectID.Hex(),
		Filename: image.Filename,
		Url: "foto url",
	}

	return responseImage, nil
}

func (r *repository) DeleteImage(ID string) error {
	imageID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.Background(), bson.M { "_id": imageID })

	if err != nil {
		return err
	}

	return nil
}
