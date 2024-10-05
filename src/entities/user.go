package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name              string
	Email             string
	Password          string
	PreviousPasswords []Password `gorm:"foreignKey:UserID"`
}
