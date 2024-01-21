package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Salon struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Name    string    `gorm:"type:varchar(255);not null"`
	Address string    `gorm:"type:varchar(255);not null"`
	Phone   string    `gorm:"type:varchar(255);not null"`
	Users   []User    `gorm:"foreignKey:SalonID"`
}
