package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
) 

type Project struct {
	ID         		     primitive.ObjectID `json:"id" bson:"_id"`
	Owner      		     primitive.ObjectID `json:"owner" bson:"owner"`
	Participants       []primitive.ObjectID `json:"participants" bson:"participants"`
	Title      		     string `json:"title" bson:"title"`
	Description 		   string `json:"description" bson:"description"`
	CreatedAt  		     time.Time `json:"createdAt" bson:"createdAt"`
	HasBeenSoftDeleted bool `json:"hasBeenSoftDeleted" bson:"hasBeenSoftDeleted"`
  Classes            []Class `json:"classes" bson:"classes"`
}

type ProjectPostRequest struct {
	Owner       string `json:"owner" bson:"owner"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type ProjectPutRequest struct {
	ID         string `json:"id" bson:"_id"`
	Owner      string `json:"owner" bson:"owner"`
	Title      string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}
