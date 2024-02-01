package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Slot struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"type:uuid;primary_key"`
	Date                time.Time `gorm:"type:date;not null"`
	BeginningHour       time.Time `gorm:"type:timestamp;not null"`
	EndTime             time.Time `gorm:"type:time;not null"`
	HairdressingStaffID uuid.UUID `gorm:"type:uuid;not null"`
	HairdressingStaff   User      `gorm:"foreignKey:HairdressingStaffID"`
}
