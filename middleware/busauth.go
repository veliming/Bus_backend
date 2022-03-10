package middleware

import (
	"Bus/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BusAUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("AuthKey")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求参数错误"})
			c.Abort()
			return
		}
		if token != setting.BusKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "鉴权失败"})
			c.Abort()
			return
		}
		c.Next()
	}
}
