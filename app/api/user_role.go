package api

import (
	"github.com/gin-gonic/gin"
	"mgo-gin/app/form"
	"mgo-gin/app/repository"
	"mgo-gin/db"
	err2 "mgo-gin/utils/err"
	"net/http"
)

func ApplyUserRoleAPI(app *gin.RouterGroup,resource *db.Resource)  {
	userRoleEntity:=repository.NewUserRoleEntity(resource)

	userRoleRoute := app.Group("/user-role")

	userRoleRoute.GET("", getAllUserRole(userRoleEntity))
	userRoleRoute.POST("", createUserRole(userRoleEntity))
}

func getAllUserRole(userRoleEntity repository.IUserRole) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		list, code, err := userRoleEntity.GetAll()
		response := map[string]interface{}{
			"userRoles": list,
			"err":  err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}

func createUserRole(userRoleEntity repository.IUserRole) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		userRoleForm := form.UserRoleForm{}
		if err := ctx.Bind(&userRoleForm); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		userRole, code, err := userRoleEntity.Create(userRoleForm)
		response := map[string]interface{}{
			"userRole": userRole,
			"err":      err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}
