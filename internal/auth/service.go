package auth

import (
	"errors"
	"github.com/GuaiZai233/Larvar/internal/db"
)

func LoginUser(username string, password string) (string, error) {
	// get DB
	database := db.GetDB()

	// query if user exists
	user, err := IsUserExist(database, username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("用户名或密码错误") // 用户不存在
	}

	// verify password
	if !CheckPswHash(password, user.PasswordHash) {
		return "", errors.New("账号或密码错误")
	}

	// generate token
	token, err := GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", errors.New("生成令牌失败")
	}

	return token, nil

}

func RegisterUser(username, password, email string) error {
	// get DB handle
	database := db.GetDB()

	exists, err := IsUserExist(database, username)
	if err != nil {
		return errors.New("数据查询失败")
	}
	if exists != nil {
		return errors.New("用户名已存在")
	}

	hashPassword, err := HashPassword(password)
	if err != nil {
		return errors.New("密码加密失败")
	}

	err = CreateUser(database, username, hashPassword, email)
	if err != nil {
		return err
	}

	return nil
}
