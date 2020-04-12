package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetToken(ctx *gin.Context) string {
	reqToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	return strings.TrimSpace(splitToken[1])
}

func GetUsername(ctx *gin.Context) string {
	reqToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	token, _ := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	claim := token.Claims.(jwt.MapClaims)
	username := claim["username"].(string)
	return username
}
