package auth

import (
  "context"
  "errors"

	"backend/internal/pkg/entities"
	"backend/internal/api/presenters"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	ReadUserByCredentials(email, password string) (*presenters.User, error)
}

type repository struct {
  collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
  return &repository {
    collection: collection,
  }
}

func (r *repository) ReadUserByCredentials(email, password string) (*presenters.User, error) {
	filter := bson.M {
		"email":    email,
		"password": password,
	}

	var user entities.User
  err := r.collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, errors.New("user was not found.")
	}

	responseUser := &presenters.User {
		ID: user.ID.Hex(),
		Name: user.Name,
		Email: user.Email,
	}

	return responseUser, nil
}

