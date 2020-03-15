package repository

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mgo-gin/app/model"
	"mgo-gin/db"
	"net/http"
)

var ToDoEntity IToDo

type toDoEntity struct{
	resource *db.Resource
	repo *mongo.Collection
}

type IToDo interface {
	GetAll() ([]model.ToDo,int ,error)
}

//func NewToDoEntity
func NewToDoEntity(resource *db.Resource) IToDo {
	toDoRepo := resource.DB.Collection("todo")
	ToDoEntity = &toDoEntity{resource: resource, repo: toDoRepo}
	return ToDoEntity
}

func (entity *toDoEntity) GetAll() ([]model.ToDo,int,error)  {
	toDoList := []model.ToDo{}
	ctx, cancel := initContext()
	defer cancel()
	cursor,err := entity.repo.Find(ctx,bson.M{})

	if err!=nil{
		logrus.Print(err)
		return []model.ToDo{},400,err
	}

	for cursor.Next(ctx){
		var todo model.ToDo
		err = cursor.Decode(&todo)
		if err!=nil{
			logrus.Print(err)
		}
		toDoList = append(toDoList,todo)
	}
	return toDoList,http.StatusOK,nil
}
