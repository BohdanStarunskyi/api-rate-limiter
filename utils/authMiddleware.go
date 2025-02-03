package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API-Key")
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "API Key required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
