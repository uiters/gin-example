package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mgo-gin/app/repository"
	"mgo-gin/utils/arrays"
	jwt2 "mgo-gin/utils/jwt"
	"net/http"
)

func GetRolesFromToken(tokenReq string) (role []string) {
	token, _ := jwt.Parse(tokenReq, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claim := token.Claims.(jwt.MapClaims)
	var roles []string
	rolesResource := claim["roles"].([]interface{})
	for _, role := range rolesResource {
		roles = append(roles, role.(string))
	}
	if len(roles) <= 0 {
		return nil
	}
	return roles
}

func RequireAuthorization(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := jwt2.GetUsername(c)
		isAccessible := false
		userRole, _, err := repository.UserRoleEntity.GetOneByNameAndRole(username, name)
		if err != nil || userRole == nil {
			notPermission(c)
			return
		}
		access := userRole.Access

		method := c.Request.Method

		if arrays.Contains(access, method) {
			isAccessible = true
		}

		if isAccessible == false {
			notPermission(c)
			return
		}
		c.Next()
	}
}

func invalidRequest(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "Invalid request, restricted endpoint"})
	c.Abort()
}

func notPermission(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "Dont have permission"})
	c.Abort()
}
