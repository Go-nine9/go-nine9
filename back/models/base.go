package models

import (
	"time"

	"github.com/Go-nine9/go-nine9/services"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Basic fields for all models
type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// Create UUID
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID, _ = services.GenerateUUID()
	return nil
}
