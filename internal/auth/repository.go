package auth

import (
	"errors"
	"gorm.io/gorm"
)

func IsUserExist(db *gorm.DB, inputUsername string) (*User, error) {
	var user User

	result := db.Where("username = ?", inputUsername).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在，这不是错误，只是空结果
		}
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(db *gorm.DB, username, password, email string) error {
	user := User{
		Username:     username,
		PasswordHash: password,
		Email:        email,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
