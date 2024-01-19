package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	Roles     string `json:"roles"`
}
