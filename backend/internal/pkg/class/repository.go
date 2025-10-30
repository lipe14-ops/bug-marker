package class

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"backend/internal/api/presenters"
	"backend/internal/pkg/entities"
)

type Repository interface {
	CreateClass(ID string, class *entities.Class) (*presenters.Class, error)
	ReadClass(projectID, classID string) (*presenters.Class, error)
  ReadClassesByProject(ID string) ([]*presenters.Class, error)
  UpdateClass(projectID string, project *entities.Class) (*presenters.Class, error)
  DeleteClass(projectID, classID string) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository {
		collection: collection,
	}
}

func (r *repository) CreateClass(ID string, class *entities.Class) (*presenters.Class, error) {
  projectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

  class.ID = primitive.NewObjectID()
  class.CreatedAt = time.Now()
  
  _, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": projectID }, bson.M { "$push": bson.M { "classes": class } })
	
	if err != nil {
		return nil, err
	}

	responseClass := &presenters.Class {
    ID: class.ID.Hex(),
    Name: class.Name,
    Description: class.Description,
    Color: class.Color,
	}

	return responseClass, err
}

func (r *repository) ReadClassesByProject(ID string) ([]*presenters.Class, error) {
	projectObjectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	filter := bson.M {
		"_id": projectObjectID,
	}

	var project entities.Project
	err = r.collection.FindOne(context.Background(), filter).Decode(&project)

	if err != nil {
		return nil, errors.New("project was not found.")
	}

  var responseClasses []*presenters.Class

  for _, class := range project.Classes {
    responseClass := &presenters.Class {
      ID: class.ID.Hex(),
      Name: class.Name,
      Description: class.Description,
      Color: class.Color,
    }

    responseClasses = append(responseClasses, responseClass)
  }

  return responseClasses, nil
}

func (r *repository) ReadClass(projectID, classID string) (*presenters.Class, error) {
	projectObjectID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return nil, err
	}

	filter := bson.M {
		"_id": projectObjectID,
	}

	var project entities.Project
	err = r.collection.FindOne(context.Background(), filter).Decode(&project)

	if err != nil {
		return nil, errors.New("project was not found.")
	}

  for _, class := range project.Classes {
    fmt.Println(class.ID.Hex())
    if class.ID.Hex() == classID {
      responseClass := &presenters.Class {
        ID: class.ID.Hex(),
        Name: class.Name,
        Description: class.Description,
        Color: class.Color,
      }
      return responseClass, nil
    }
  }

  return nil, errors.New("class was not found.")
}

func (r *repository) UpdateClass(projectID string, class *entities.Class) (*presenters.Class, error) {
	projectObjectID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return nil, err
	}

  _, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": projectObjectID, "classes._id": class.ID }, bson.M { "$set": bson.M { "classes.$": class }})
	
	if err != nil {
		return nil, err
	}

	responseClass := &presenters.Class {
    ID: class.ID.Hex(),
    Name: class.Name,
    Description: class.Description,
    Color: class.Color,
	}

	return responseClass, nil
}

func (r *repository) DeleteClass(projectID, classID string) error {
	projectObjectID, err := primitive.ObjectIDFromHex(projectID)

	if err != nil {
		return err
	}

	classObjectID, err := primitive.ObjectIDFromHex(classID)

	if err != nil {
		return err
	}

  _, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": projectObjectID }, bson.M { "$pull": bson.M { "classes": bson.M { "_id": classObjectID }}})

	if err != nil {
		return err
	}

	return nil
}
