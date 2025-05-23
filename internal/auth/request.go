package auth

type RegisterRequest struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type LoginRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
