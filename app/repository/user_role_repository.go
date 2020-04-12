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
	"mgo-gin/utils/constant"
	"net/http"
)

var UserRoleEntity IUserRole

type userRoleEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
}

type IUserRole interface {
	GetAll() ([]model.UserRole, int, error)
	Create(form form.UserRoleForm) (model.UserRole, int, error)
	GetOneByNameAndRole(username, name string) (*model.UserRole, int, error) // need return pointer
}

func NewUserRoleEntity(resource *db.Resource) IUserRole {
	userRoleRepo := resource.DB.Collection("user_role")
	UserRoleEntity = &userRoleEntity{resource: resource, repo: userRoleRepo}
	return UserRoleEntity
}

func (entity *userRoleEntity) GetAll() ([]model.UserRole, int, error) {
	userRoleList := []model.UserRole{}
	ctx, cancel := initContext()
	defer cancel()
	cursor, err := entity.repo.Find(ctx, bson.M{})

	if err != nil {
		return []model.UserRole{}, http.StatusBadRequest, err
	}

	for cursor.Next(ctx) {
		var userRole model.UserRole
		err = cursor.Decode(&userRole)
		if err != nil {
			logrus.Print(err)
		}
		userRoleList = append(userRoleList, userRole)
	}
	return userRoleList, http.StatusOK, nil
}

func (entity *userRoleEntity) GetOneByNameAndRole(username, name string) (*model.UserRole, int, error) {
	ctx, cancel := initContext()
	defer cancel()
	var role model.UserRole

	err := entity.repo.FindOne(ctx, bson.M{"role": name, "username": username}).Decode(&role)

	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return &role, http.StatusOK, nil
}

func (entity *userRoleEntity) Create(form form.UserRoleForm) (model.UserRole, int, error) {

	ctx, cancel := initContext()
	defer cancel()

	for _, access := range form.Access {
		if access != constant.GET && access != constant.POST && access != constant.PUT && access != constant.DELETE {
			return model.UserRole{}, http.StatusBadRequest, errors.New("access is invalid")
		}
	}
	role, _, err := entity.GetOneByNameAndRole(form.Username, form.Role)
	if err != nil || role == nil {
		return model.UserRole{}, http.StatusBadRequest, err
	}

	newRole := model.UserRole{
		Id:     primitive.NewObjectID(),
		Role:   form.Role,
		Access: form.Access,
	}
	_, err = entity.repo.InsertOne(ctx, newRole)
	if err != nil {
		return model.UserRole{}, http.StatusBadRequest, err
	}
	return newRole, http.StatusOK, nil
}
