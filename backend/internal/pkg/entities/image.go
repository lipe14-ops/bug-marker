package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
) 

type Image struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	ProjectID  primitive.ObjectID `json:"projectID" bson:"projectID"`
	Filename	 string `json:"filename" bson:"filename"`
  CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
  HasBeenSoftDeleted bool `json:"hasBeenSoftDeleted" bson:"hasBeenSoftDeleted"`
  Polygons   []Polygon `json:"polygons" bson:"polygons"`
}

type ImagePostRequest struct {
  ID         string `json:"id" bson:"_id"`
  ProjectID  string `json:"projectID" bson:"projectID"`
	Filename	 string `json:"filename" bson:"filename"`
}

type ImagePutRequest struct {
	Filename	 string `json:"filename" bson:"filename"`
}
