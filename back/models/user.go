package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	LastName  string     `gorm:"type:varchar(255);not null"`
	FirstName string     `gorm:"type:varchar(255);not null"`
	Email     string     `gorm:"type:varchar(255);not null"`
	Password  string     `gorm:"type:varchar(255);not null"`
	Roles     string     `gorm:"type:varchar(255);not null"`
	SalonID   *uuid.UUID `gorm:"type:uuid"`
}
