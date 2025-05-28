package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loiclaborderie/bahasa-project/constants"
	"github.com/loiclaborderie/bahasa-project/pkg/helper"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := helper.ValidateTokenAndGetUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func RequiresRole(roleRequired constants.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helper.ValidateTokenAndGetUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.Role != roleRequired {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "you don't have the permissions for that"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
