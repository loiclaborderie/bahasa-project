package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/loiclaborderie/bahasa-project/constants"
	"github.com/loiclaborderie/bahasa-project/internal/user"
	"github.com/loiclaborderie/bahasa-project/pkg/db"
)

// Helper function to handle the token validation and user retrieval
func validateTokenAndGetUser(c *gin.Context) (*user.User, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("you don't have the rights")
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		return nil, fmt.Errorf("invalid token format")
	}

	tokenStr := authToken[1]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, fmt.Errorf("token expired")
	}

	var user user.User
	if err := db.GetDB().Where("email = ?", claims["email"]).Find(&user).Error; err != nil || user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := validateTokenAndGetUser(c)
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
		user, err := validateTokenAndGetUser(c)
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
