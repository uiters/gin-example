package api

import (
	"github.com/gin-gonic/gin"
	"mgo-gin/app/form"
	"mgo-gin/app/repository"
	"mgo-gin/db"
	err2 "mgo-gin/utils/err"
	"net/http"
)

func ApplyRoleAPI(app *gin.RouterGroup,resource *db.Resource)  {
	roleEntity:=repository.NewRoleEntity(resource)

	roleRoute := app.Group("/roles")

	roleRoute.GET("", getAllRoles(roleEntity))
	roleRoute.POST("", createRole(roleEntity))
}

func getAllRoles(roleEntity repository.IRole) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		list, code, err := roleEntity.GetAll()
		response := map[string]interface{}{
			"roles": list,
			"err":  err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}

func createRole(roleEntity repository.IRole) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		roleForm := form.RoleForm{}
		if err := ctx.Bind(&roleForm); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		role, code, err := roleEntity.CreateOne(roleForm)
		response := map[string]interface{}{
			"role": role,
			"err":  err2.GetErrorMessage(err),
		}
		ctx.JSON(code, response)
	}
}
