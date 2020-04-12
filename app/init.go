package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mgo-gin/app/api"
	"mgo-gin/app/repository"
	"mgo-gin/db"
	"mgo-gin/middlewares"
	"mgo-gin/utils/constant"
	"os"
)

type Routes struct {
}

func (app Routes) StartGin() {
	r := gin.Default()
	publicRoute := r.Group("/api/v1")
	resource, err := db.InitResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	r.Use(gin.Logger())
	r.Use(middlewares.NewRecovery())
	r.Use(middlewares.NewCors([]string{"*"}))
	r.GET("swagger/*any", middlewares.NewSwagger())

	r.Static("/template/css", "./template/css")
	r.Static("/template/images", "./template/images")
	//r.Static("/template", "./template")

	r.NoRoute(func(context *gin.Context) {
		//context.File("./template/route_not_found.html")
		context.File("./template/index.html")
	})
	initAllAPI()
	api.ApplyToDoAPI(publicRoute, resource)
	api.ApplyRoleAPI(publicRoute, resource)
	api.ApplyUserRoleAPI(publicRoute, resource)
	api.ApplyUserAPI(publicRoute, resource)
	repository.UserEntity.Init()
	r.Run(":" + os.Getenv("PORT"))
}

func initAllAPI(){
	constant.Controller["TODO"]= "TODO"
	constant.Controller["ROLE"]= "ROLE"
	constant.Controller["USER"]= "USER"
	constant.Controller["USER_ROLE"]= "USER_ROLE"
}
