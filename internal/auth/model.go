package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"unique"`
}
