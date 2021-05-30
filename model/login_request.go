package model

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password"`
	RememberMe bool `json:"remember_me"`
}
