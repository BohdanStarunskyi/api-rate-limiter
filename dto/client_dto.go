package dto

type CreateClientRequest struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	RateLimit int    `json:"rate_limit" binding:"required"`
	Password  string `json:"password" binding:"required,min=8,max=128"`
}

type LoginClientRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
