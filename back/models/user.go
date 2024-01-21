package models

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
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

func VerifyPassword(hashedPassword string, password string) error {
	if strings.HasPrefix(hashedPassword, "$2a$") {
		return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	}
	if hashedPassword != password {
		return errors.New("mot de passe incorrect")
	}
	return nil
}
