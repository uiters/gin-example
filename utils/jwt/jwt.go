package jwt

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetToken(ctx *gin.Context) string {
	reqToken := ctx.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	return strings.TrimSpace(splitToken[1])
}
