package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")
		if userRole != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this page."})
			c.Abort()
			return
		}
		c.Next()
	}
}
