package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mgo-gin/app/api"
	"mgo-gin/db"
)

type Routes struct{

}

func (app Routes) StartGin() {
	r := gin.Default()
	publicRoute :=r.Group("/api/v1")
	resource, err:= db.InitResource()
	if err!=nil{
		logrus.Print(err)
	}
	api.ApplyToDoAPI(publicRoute,resource)
	r.Run(":8585")
}
