package repository

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mgo-gin/app/form"
	"mgo-gin/app/model"
	"mgo-gin/db"
	"net/http"
)

var UserEntity IUser

type userEntity struct{
	resource *db.Resource
	repo *mongo.Collection
}

type IUser interface {
	GetAll() ([]model.User,int ,error)
	GetOneByUsernameAndPassword(userForm form.User) (model.User,int ,error)
}

//func NewToDoEntity
func NewUserEntity(resource *db.Resource) IUser {
	userRepo := resource.DB.Collection("user")
	UserEntity = &userEntity{resource: resource, repo: userRepo}
	return UserEntity
}

func (entity *userEntity) GetAll() ([]model.User,int,error)  {
	usersList :=[]model.User{}
	ctx, cancel := initContext()
	defer cancel()
	cursor,err := entity.repo.Find(ctx,bson.M{})

	if err!=nil{
		logrus.Print(err)
		return []model.User{},400,err
	}

	for cursor.Next(ctx){
		var user model.User
		err = cursor.Decode(&user)
		if err!=nil{
			logrus.Print(err)
		}
		usersList = append(usersList, user)
	}
	return usersList,http.StatusOK,nil
}

func (entity *userEntity) GetOneByUsernameAndPassword(userForm form.User) (model.User,int,error)  {
	ctx, cancel := initContext()
	defer cancel()

	objID,err :=primitive.ObjectIDFromHex(userForm.Username)
	if err!=nil{
		logrus.Print(err)
	}

	var user model.User
	err =
		entity.repo.FindOne(ctx,bson.M{"_id" : objID, "password": userForm.Password},).Decode(&user)

	if err!=nil{
		logrus.Print(err)
		return model.User{},400,err
	}

	return user,http.StatusOK,nil
}
