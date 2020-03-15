package api

import (
	"github.com/gin-gonic/gin"
	"mgo-gin/app/form"
	"mgo-gin/app/repository"
	"mgo-gin/db"
	"mgo-gin/middlewares"
	err2 "mgo-gin/utils/err"
	"net/http"
)

func ApplyUserAPI(app *gin.RouterGroup, resource *db.Resource) {
	userEntity := repository.NewUserEntity(resource)
	authRoute := app.Group("")
	authRoute.POST("/login", login(userEntity))
	authRoute.POST("/sign-up", signUp(userEntity))

	userRoute := app.Group("/user")
	userRoute.GET("", getAllUSer(userEntity))
}

func login(userEntity repository.IUser) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		userRequest := form.User{}
		if err := ctx.Bind(&userRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		user, code, _ := userEntity.GetOneByUsernameAndPassword(userRequest)
		token := middlewares.GenerateJWTToken(user)
		response := map[string]interface{}{
			"token": token,
			"error": nil,
		}
		ctx.JSON(code, response)
	}
}

func signUp(userEntity repository.IUser) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		userRequest := form.User{}
		if err := ctx.Bind(&userRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		user, code, err := userEntity.CreateOne(userRequest)
		response := map[string]interface{}{
			"user": user,
			"error":  err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}

func getAllUSer(userEntity repository.IUser) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		list, code, err := userEntity.GetAll()
		response := map[string]interface{}{
			"users": list,
			"error":   err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}
