package api

import (
	"github.com/gin-gonic/gin"
	"mgo-gin/app/form"
	"mgo-gin/app/repository"
	"mgo-gin/db"
)

func ApplyUserAPI(app *gin.RouterGroup, resource *db.Resource){
	userEntity := repository.NewUserEntity(resource)
	authRoute := app.Group("")
	authRoute.POST("", login(userEntity))

	userRoute := app.Group("/user")
	userRoute.GET("", getAllUSer(userEntity))
}

func login(userEntity repository.IUser) func (ctx *gin.Context){
	return func (ctx *gin.Context){

		userRequest := form.User{}
		user,code,err:=userEntity.GetOneByUsernameAndPassword(userRequest)
		response := map[string]interface{}{
			"user":  user,
			"error": err,
		}
		ctx.JSON(code,response)
	}
}

func getAllUSer(userEntity repository.IUser) func (ctx *gin.Context){
	return func (ctx *gin.Context){
		list,code,err:= userEntity.GetAll()
		response := map[string]interface{}{
			"users":  list,
			"error": err,
		}
		ctx.JSON(code,response)
	}
}
