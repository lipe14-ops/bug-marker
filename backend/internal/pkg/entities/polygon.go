package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
) 

type Polygon struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	ImageID         primitive.ObjectID `json:"imageID" bson:"imageID"`
	ClassID         primitive.ObjectID `json:"classID" bson:"classID"`
	Name            string `json:"name" bson:"name"`
	Type            string `json:"type" bson:"type"`
	Coordinates     [][][]float64 `json:"coordinates" bson:"coordinates"`
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
}

type PolygonPostRequest struct {
	Name            string `json:"name" bson:"name"`
	ClassID         string `json:"classID" bson:"classID"`
	Coordinates     [][]float64 `json:"coordinates" bson:"coordinates"`
}

type PolygonPutRequest struct {
	Name            string `json:"name" bson:"name"`
	ClassID         string `json:"classID" bson:"classID"`
	Coordinates     [][]float64 `json:"coordinates" bson:"coordinates"`
}




