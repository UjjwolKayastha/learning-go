package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
	Base
	ID          string `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

// TableName of the model
func (h *Hotel) TableName() string {
	return "hotels"
}

// BeforeCreate -> add uuid before create
func (h *Hotel) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	h.ID = id.String()
	return err
}
