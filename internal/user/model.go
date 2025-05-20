package user

import (
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"-"` // Password is not returned in JSON
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
