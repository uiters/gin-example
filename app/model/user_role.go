package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRole struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `bson:"username" json:"username"`
	Role     string             `bson:"role" json:"role"`
	Access   []string           `bson:"access" json:"access"`
}
