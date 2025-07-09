package middleware

import (
	"net/http"
	"strings"

	"github.com/OgiDac/CompanyTask/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 {
			authToken := parts[1]
			authorized, err := utils.IsAuthorized(authToken, secret)
			if err != nil || !authorized {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}

			userID, err := utils.ExtractIDFromToken(authToken, secret)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			// Store user ID in context
			c.Set("user_id", userID)
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
	}
}
