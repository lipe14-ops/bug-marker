package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
) 

type Project struct {
	ID         		primitive.ObjectID `json:"id" bson:"_id"`
	Owner      		primitive.ObjectID `json:"owner" bson:"owner"`
	Participants  []primitive.ObjectID `json:"participants" bson:"participants"`
	Title      		string `json:"title" bson:"title"`
	Description 		string `json:"description" bson:"description"`
	CreatedAt  		time.Time `json:"createdAt" bson:"createdAt"`
	HasBeenSoftDeleted bool `json:"hasBeenSoftDeleted" bson:"hasBeenSoftDeleted"`
}

type ProjectPostRequest struct {
	Owner       string `json:"owner" bson:"owner"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type ProjectPutRequest struct {
	Owner      string `json:"owner" bson:"owner"`
	Title      string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}
