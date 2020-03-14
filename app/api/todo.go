package api

import (
	"github.com/gin-gonic/gin"
	"mgo-gin/app/repository"
	"mgo-gin/db"
)

func ApplyToDoAPI(app *gin.RouterGroup, resource *db.Resource){
	toDoEntity := repository.NewToDoEntity(resource)
	toDoRoute := app.Group("/todo")
	toDoRoute.GET("", getAllToDo(toDoEntity))

}

func getAllToDo(toDoEntity repository.IToDo) func (ctx *gin.Context){
	return func (ctx *gin.Context){
		list,code,err:=toDoEntity.GetAll()
		response := map[string]interface{}{
			"todo":  list,
			"error": err,
		}
		ctx.JSON(code,response)
	}
}