package util

import (
	"github.com/gin-gonic/gin"
	"log"
)

func CheckRole(ctx *gin.Context, role int) (*Claims, bool) {
	temp, exists := ctx.Get("claims")
	if !exists {
		return nil, false
	}
	claims, r := temp.(*Claims)
	if !r {
		log.Println("Faild get claims at ", ctx.ClientIP())
		ctx.Abort()
	}
	return claims, claims.Role == role
}
