package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"-"` // Password is not returned in JSON
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser creates a new user with default values
func NewUser(username, email, password string) *User {
	now := time.Now()
	return &User{
		ID:        generateID(),
		Username:  username,
		Email:     email,
		Password:  password, // Note: In production, this should be hashed
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// generateID creates a unique ID for a user
// In a real application, you might use a UUID library
func generateID() string {
	return time.Now().Format("20060102150405")
}