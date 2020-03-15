package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	id       primitive.ObjectID `bson:"id" json:"id"`
	username string             `bson:"username" json:"username"`
	password string             `bson:"password" json:"password"`
	roles    []string           `bson:"roles" json:"roles"`
}
