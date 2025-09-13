package s2s

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware(validToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing service token"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != validToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid service token"})
			return
		}

		c.Next()
	}
}
