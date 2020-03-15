package api

import (
	"github.com/gin-gonic/gin"
	"mgo-gin/app/form"
	"mgo-gin/app/repository"
	"mgo-gin/db"
	err2 "mgo-gin/utils/err"
	"net/http"
)

func ApplyToDoAPI(app *gin.RouterGroup, resource *db.Resource) {
	toDoEntity := repository.NewToDoEntity(resource)
	toDoRoute := app.Group("/todo")

	toDoRoute.GET("", getAllToDo(toDoEntity))
	toDoRoute.GET("/:id", getToDoById(toDoEntity))
	toDoRoute.POST("", createToDo(toDoEntity))
	toDoRoute.PUT("/:id", updateToDo(toDoEntity))
}

func getAllToDo(toDoEntity repository.IToDo) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		list, code, err := toDoEntity.GetAll()
		response := map[string]interface{}{
			"todo": list,
			"err":  err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}

func createToDo(toDoEntity repository.IToDo) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		todoReq := form.ToDoForm{}
		if err := ctx.Bind(&todoReq); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		todo, code, err := toDoEntity.CreateOne(todoReq)
		response := map[string]interface{}{
			"todo": todo,
			"err":  err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}

func getToDoById(toDoEntity repository.IToDo) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id :=ctx.Param("id")
		todo, code, err := toDoEntity.GetOneByID(id)
		response := map[string]interface{}{
			"todo": todo,
			"err":  err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}

func updateToDo(toDoEntity repository.IToDo) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		id :=ctx.Param("id")
		todoReq := form.ToDoForm{}
		if err := ctx.Bind(&todoReq); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		todo, code, err := toDoEntity.Update(id,todoReq)
		response := map[string]interface{}{
			"todo": todo,
			"err":  err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}
