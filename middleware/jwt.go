package middleware

import (
	"Bus/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求参数错误"})
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token鉴权失败"})
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token已超时"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
