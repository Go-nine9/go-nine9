package models

import "github.com/google/uuid"

type Salon struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Name    string    `gorm:"type:varchar(255);not null"`
	Address string    `gorm:"type:varchar(255);not null"`
	Phone   string    `gorm:"type:varchar(255);not null"`
}
