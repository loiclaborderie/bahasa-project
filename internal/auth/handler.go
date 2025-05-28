package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/loiclaborderie/bahasa-project/internal/user"
	"github.com/loiclaborderie/bahasa-project/pkg/helper"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Db       *gorm.DB
	Validate *validator.Validate
}

func NewAuthHandlerImpl(Db *gorm.DB, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{Db: Db, Validate: validate}
}

func (c AuthHandler) Register(ctx *gin.Context) {
	var reqBody RegisterRequest

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errMsg := fmt.Sprintf("Validation failed for field: %s", validationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var existingUser user.User

	result := c.Db.Where("email = ?", reqBody.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	password, err := helper.EncryptPassword(reqBody.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong, failed to encrypt password"})
		return
	}

	newUser := user.User{
		Email:    reqBody.Email,
		Password: password,
		Username: reqBody.Username,
	}

	if err := c.Db.Create(&newUser).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	if err := setLoginCookie(ctx, newUser.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "JWT Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": existingUser.ToUserResponse(), "message": "User registered and logged in successfully"})
}

func (c AuthHandler) Login(ctx *gin.Context) {
	var reqBody LoginRequest

	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errMsg := fmt.Sprintf("Validation failed for field: %s", validationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	var existingUser user.User
	result := c.Db.Where("email = ?", reqBody.Email).First(&existingUser)
	if result.RowsAffected < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "There is no user registered with this email"})
		return
	}

	valid := helper.ComparePassword(reqBody.Password, existingUser.Password)

	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password invalid"})
		return
	}

	err := setLoginCookie(ctx, existingUser.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "JWT Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "user": existingUser.ToUserResponse()})
}

func (c AuthHandler) Logout(ctx *gin.Context) {
	_, err := ctx.Cookie("token")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "You are not logged in yet"})
		return
	}

	ctx.SetCookie(
		"token", // name
		"",      // empty value
		-1,      // maxAge (negative to expire immediately)
		"/",     // path
		"",      // domain
		false,   // secure
		true,    // httpOnly
	)

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (c AuthHandler) Me(ctx *gin.Context) {
	user, err := helper.ValidateTokenAndGetUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user.ToUserResponse()})
}

func setLoginCookie(ctx *gin.Context, email string) error {
	token, err := helper.CreateToken(email)
	if err != nil {
		return err
	}
	ctx.SetCookie("token", token, int(24*time.Hour), "/", "", false, true)
	return nil
}
