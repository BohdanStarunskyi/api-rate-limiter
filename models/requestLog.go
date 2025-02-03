package models

import (
	"time"
)

type RequestLog struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	ClientID  uint      `json:"client_id"`
	Endpoint  string    `json:"endpoint"`
	Timestamp time.Time `json:"timestamp"`
}
