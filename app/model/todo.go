package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	Id   primitive.ObjectID `bson:"id" json:"id"`
	Name string             `bson:"name" json:"name"`
}
