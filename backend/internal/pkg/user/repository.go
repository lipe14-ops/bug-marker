package user

import (
	"time"
	"errors"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"

	"backend/internal/pkg/entities"
	"backend/internal/api/presenters"
)

type Repository interface {
	CreateUser(user *entities.User) (*presenters.User, error)
	ReadUser(ID string) (*presenters.User, error)
	UpdateUser(ID string, user *entities.User) (*presenters.User, error)
	DeleteUser(ID string) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return &repository {
		collection: collection,
	}
}

func (r *repository) CreateUser(user *entities.User) (*presenters.User, error) {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(context.Background(), user)

	if err != nil {
		return nil, err
	}

	responseUser := &presenters.User {
		ID: user.ID.Hex(),
		Name: user.Name,
		Email: user.Email,
	}

	return responseUser, err
}

func (r *repository) ReadUser(ID string) (*presenters.User, error) {
	userID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	filter := bson.M {
		"_id": userID,
	}

	var user entities.User

	err = r.collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return nil, errors.New("user was not found.")
	}

	responseUser := &presenters.User {
		ID: ID,
		Name: user.Name,
		Email: user.Email,
	}

	return responseUser, nil
}

func (r *repository) UpdateUser(ID string, user *entities.User) (*presenters.User, error) {
	userID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return nil, err
	}

	user.ID = userID
	
	_, err = r.collection.UpdateOne(context.Background(), bson.M { "_id": userID }, bson.M { "$set": user })
	
	if err != nil {
		return nil, err
	}

	responseUser := &presenters.User {
		ID: ID,
		Name: user.Name,
		Email: user.Email,
	}

	return responseUser, nil
}

func (r *repository) DeleteUser(ID string) error {
	userID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.Background(), bson.M { "_id": userID })

	if err != nil {
		return err
	}

	return nil
}
