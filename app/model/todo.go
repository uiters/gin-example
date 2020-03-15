package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	id   primitive.ObjectID `bson:"id" json:"id"`
	name string             `bson:"name" json:"name"`
}
