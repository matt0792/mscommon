package s2s

import (
	"context"
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

		tenantID := c.GetHeader("X-Internal-Tenant-ID")
		userID := c.GetHeader("X-Internal-User-ID")
		roles := strings.Split(c.GetHeader("X-Internal-Roles"), ",")

		ctx := context.WithValue(c.Request.Context(), S2SContextKey{}, S2SRequestContext{
			TenantID: tenantID,
			UserID:   userID,
			Roles:    roles,
		})

		c.Request = c.Request.WithContext(ctx)
		c.Set("s2s_context", ctx)

		c.Next()
	}
}
