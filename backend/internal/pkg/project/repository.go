package project

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
	CreateProject(ID string, project *entities.Project) (*presenters.Project, error)
	ReadProject(ID string) (*presenters.Project, error)
	ReadProjectByOwner(ID string) ([]*presenters.Project, error)
	UpdateProject(ID string, project *entities.Project) (*presenters.Project, error)
	DeleteProject(ID string) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository {
		collection: collection,
	}
}

func (r *repository) CreateProject(ID string, project *entities.Project) (*presenters.Project, error) {
  ownerID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	project.ID = primitive.NewObjectID()
	project.Owner = ownerID
	project.CreatedAt = time.Now()
	project.Participants = []primitive.ObjectID{}
  project.Classes      = []entities.Class{}
  project.HasBeenSoftDeleted = false

	_, err = r.collection.InsertOne(context.Background(), project)

	if err != nil {
		return nil, err
	}

	responseProject := &presenters.Project {
		ID: project.ID.Hex(),
		Owner: ID,
		Title: project.Title,
		Description: project.Description,
		Participants: []string{},
    HasBeenSoftDeleted: project.HasBeenSoftDeleted,
	}

	return responseProject, err
}

func (r *repository) ReadProjectByOwner(ID string) ([]*presenters.Project, error) {
	ownerID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	filter := bson.M {
		"owner": ownerID,
	}

  cursor, err := r.collection.Find(context.TODO(), filter)

	var projects []entities.Project

  if err = cursor.All(context.TODO(), &projects); err != nil {
    return nil, err
  }

  var responseProjects []*presenters.Project
  
  for _, project := range projects {
    responseProject := &presenters.Project {
      ID: project.ID.Hex(),
      Owner: project.Owner.Hex(),
      Title: project.Title,
      Description: project.Description,
      Participants: []string {},
      HasBeenSoftDeleted: project.HasBeenSoftDeleted,
    }

    for _, participantID := range project.Participants {
      responseProject.Participants = append(responseProject.Participants, participantID.Hex())
    }

    responseProjects = append(responseProjects, responseProject)
  }

	return responseProjects, nil
}

func (r *repository) ReadProject(ID string) (*presenters.Project, error) {
	projectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	filter := bson.M {
		"_id": projectID,
	}

	var project entities.Project
	err = r.collection.FindOne(context.Background(), filter).Decode(&project)

	if err != nil {
		return nil, errors.New("project was not found.")
	}

	responseProject := &presenters.Project {
		ID: project.ID.Hex(),
		Owner: project.Owner.Hex(),
		Title: project.Title,
		Description: project.Description,
    Participants: []string {},
    HasBeenSoftDeleted: project.HasBeenSoftDeleted,
	}

	for _, participantID := range project.Participants {
		responseProject.Participants = append(responseProject.Participants, participantID.Hex())
	}

	return responseProject, nil
}

func (r *repository) UpdateProject(ID string, project *entities.Project) (*presenters.Project, error) {
	projectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	project.ID = projectID
	
	_, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": projectID }, bson.M { "$set": project })
	
	if err != nil {
		return nil, err
	}

	responseProject := &presenters.Project {
		ID: project.ID.Hex(),
		Owner: project.Owner.Hex(),
		Title: project.Title,
		Description: project.Description,
    HasBeenSoftDeleted: project.HasBeenSoftDeleted,
	}

	for _, participantID := range project.Participants {
		responseProject.Participants = append(responseProject.Participants, participantID.Hex())
	}

	return responseProject, nil
}

func (r *repository) DeleteProject(ID string) error {
	projectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	//_, err = r.collection.DeleteOne(context.Background(), bson.M { "_id": projectID })
	_, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": projectID }, bson.M { "$set": bson.M { "hasBeenSoftDeleted": true } })

	if err != nil {
		return err
	}

	return nil
}
