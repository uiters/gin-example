package middlewares

import (
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
)

//func NewRecovery
func NewRecovery() gin.HandlerFunc {
	return nice.Recovery(recoveryHandler)
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.HTML(500, "error.tmpl", gin.H{
		"title": "Error",
		"error": err,
	})
}
