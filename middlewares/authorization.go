package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
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

func RequireAuthorization(auths ...string)  gin.HandlerFunc {
	return func(c *gin.Context) {
		token:=c.GetHeader("Authorization")
		if token ==""{
			logrus.Print("abc")
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		jwtToken :=strings.Split(token,"Bearer ")
		roles := GetRolesFromToken(jwtToken[1])
		if len(roles) <= 0 {
			invalidRequest(c)
			return
		}
		isAccessible := false
		if len(roles) < len(auths) || len(roles) == len(auths) {
			for _, auth := range auths {
				for _, role := range roles {
					if role == auth {
						isAccessible = true
						break
					}
				}
			}
		}
		if len(roles) > len(auths) {
			for _, role := range roles {
				for _, auth := range auths {
					if auth == role {
						isAccessible = true
						break
					}
				}
			}
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
