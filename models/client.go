package models

import "time"

type Client struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	APIKey    string    `json:"api_key"`
	RateLimit int       `json:"rate_limit"`
	CreatedAt time.Time `json:"created_at"`
}
