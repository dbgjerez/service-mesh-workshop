package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Film struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Duration int                `bson:"duration,omitempty"`
	Premium  bool               `bson:"premium,omitempty"`
}
