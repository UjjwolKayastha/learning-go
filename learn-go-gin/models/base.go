package models

import (
	"time"

	"gorm.io/gorm"
)

// Base base model
type Base struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"udpated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
