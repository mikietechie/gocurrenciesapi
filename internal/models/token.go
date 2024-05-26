package models

import (
	"time"

	"gorm.io/gorm"
)

type BlackToken struct {
	gorm.Model
	Token     string `json:"token"`
	ExpiresAt time.Time
}
