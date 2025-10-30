package entities

import (
  "time"
	"go.mongodb.org/mongo-driver/bson/primitive"
) 

type Class struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name   		  string `json:"name" bson:"name"`
	Color       string `json:"color" bson:"color"`
	Description string `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

type ClassPostRequest struct {
	Name   		  string `json:"name" bson:"name"`
	Color       string `json:"color" bson:"color"`
	Description string `json:"description" bson:"description"`
}

type ClassPutRequest struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name   		  string `json:"name" bson:"name"`
	Color       string `json:"color" bson:"color"`
	Description string `json:"description" bson:"description"`
}
