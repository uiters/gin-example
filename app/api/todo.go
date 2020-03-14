package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApplyToDoAPI(app *gin.RouterGroup){
	toDoRoute := app.Group("/todo")
	toDoRoute.GET("",getAllToDo())

}

func getAllToDo() func (ctx *gin.Context){
	return func (ctx *gin.Context){
		response := map[string]interface{}{
			"test":  "ABC",
			"error": "",
		}
		ctx.JSON(http.StatusOK,response)
	}
}