package models

import (
	"time"

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

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.ID = uuid
	return nil
}
