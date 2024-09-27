package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PrimeVideo struct {
	MovieID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieName string             `json:"name,omitempty"`
	Genre     string             `json:"genre,omitempty"`
	Watched   bool               `json:"watched,omitempty"`
}
