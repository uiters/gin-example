package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mgo-gin/app/form"
	"mgo-gin/app/model"
	"mgo-gin/db"
	"net/http"
	"strings"
)

var RoleEntity IRole

type roleEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
}

type IRole interface {
	GetAll() ([]model.Role, int, error)
	CreateOne(roleForm form.RoleForm) (model.Role, int, error)
	GetOneByID(id string) (*model.Role, int, error)     // need return pointer
	GetOneByName(name string) (*model.Role, int, error) // need return pointer
}

//func NewToDoEntity
func NewRoleEntity(resource *db.Resource) IRole {
	roleRepo := resource.DB.Collection("role")
	RoleEntity = &roleEntity{resource: resource, repo: roleRepo}
	return RoleEntity
}

func (entity *roleEntity) GetAll() ([]model.Role, int, error) {
	rolesList := []model.Role{}
	ctx, cancel := initContext()
	defer cancel()
	cursor, err := entity.repo.Find(ctx, bson.M{})

	if err != nil {
		return []model.Role{}, http.StatusBadRequest, err
	}

	for cursor.Next(ctx) {
		var role model.Role
		err = cursor.Decode(&role)
		if err != nil {
			logrus.Print(err)
		}
		rolesList = append(rolesList, role)
	}
	return rolesList, http.StatusOK, nil
}

func (entity *roleEntity) GetOneByName(name string) (*model.Role, int, error) {
	ctx, cancel := initContext()
	defer cancel()
	var role model.Role

	err := entity.repo.FindOne(ctx, bson.M{"name": name}).Decode(&role)

	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return &role, http.StatusOK, nil
}

func (entity *roleEntity) GetOneByID(id string) (*model.Role, int, error) {
	ctx, cancel := initContext()
	defer cancel()
	var role model.Role
	objID, _ := primitive.ObjectIDFromHex(id)
	err := entity.repo.FindOne(ctx, bson.M{"_id": objID}).Decode(&role)

	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return &role, http.StatusOK, nil
}

func (entity *roleEntity) CreateOne(roleForm form.RoleForm) (model.Role, int, error) {

	ctx, cancel := initContext()
	defer cancel()

	role, _, err := entity.GetOneByName(strings.ToUpper(roleForm.Name))
	if err == nil || role != nil {
		return model.Role{}, http.StatusBadRequest, errors.New("this role is exist")
	}

	newRole:=model.Role{
		Id:   primitive.NewObjectID(),
		Name: strings.ToUpper(roleForm.Name),
	}
	_,err= entity.repo.InsertOne(ctx,newRole)
	if err != nil {
		return model.Role{}, http.StatusBadRequest, err
	}
	return newRole, http.StatusOK, nil
}
