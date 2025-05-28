package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/loiclaborderie/bahasa-project/internal/auth"
	"github.com/loiclaborderie/bahasa-project/middleware"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(db *gorm.DB, router *gin.Engine) {
	auth := auth.NewAuthHandlerImpl(db, validator.New())
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
	router.POST("/logout", middleware.RequireAuth(), auth.Logout)
	router.GET("/me", middleware.RequireAuth(), auth.Me)
}
