package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Lastname  string     `gorm:"type:varchar(255);not null"`
	Firstname string     `gorm:"type:varchar(255);not null"`
	Email     string     `gorm:"type:varchar(255);not null"`
	Password  string     `gorm:"type:varchar(255);not null"`
	Roles     string     `gorm:"type:varchar(255);not null"`
	SalonID   *uuid.UUID `gorm:"type:uuid"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
