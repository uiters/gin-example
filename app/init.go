package app

import (
	"github.com/gin-gonic/gin"
	"mgo-gin/app/api"
	"mgo-gin/db"
)

type Routes struct{

}

func (app Routes) StartGin() {
	r := gin.Default()
	publicRoute :=r.Group("/api/v1")
	db.InitResource()
	api.ApplyToDoAPI(publicRoute)
	r.Run(":8585")
}
