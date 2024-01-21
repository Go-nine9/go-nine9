package models

type Salon struct {
	Base
	Name    string `gorm:"type:varchar(255);not null"`
	Address string `gorm:"type:varchar(255);not null"`
	Phone   string `gorm:"type:varchar(255);not null"`
}
