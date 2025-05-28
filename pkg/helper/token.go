package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/loiclaborderie/bahasa-project/internal/user"
	"github.com/loiclaborderie/bahasa-project/pkg/db"
)

func CreateToken(email string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
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

	return claims, nil
}

func ValidateTokenAndGetUser(c *gin.Context) (*user.User, error) {
	tokenString, cookieErr := c.Cookie("token")

	if cookieErr != nil {
		return nil, fmt.Errorf("no auth token was found")
	}

	claims, err := ValidateToken(tokenString)

	if err != nil {
		return nil, err
	}

	var user user.User
	if err := db.GetDB().Where("email = ?", claims["email"]).Find(&user).Error; err != nil || user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}
