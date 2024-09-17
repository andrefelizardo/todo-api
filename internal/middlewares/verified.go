package middlewares

import (
	"net/http"

	"github.com/andrefelizardo/todo-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func Authorize(requiredVerified bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Claims not found"})
			return
		}

		userClaims := claims.(*utils.Claims)

		if requiredVerified && !userClaims.IsVerified {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Email not verified"})
			return
		}

		c.Next()
	}
}