package utils

import (
	"golang.org/x/crypto/bcrypt"
)

//
func UserPassWordCryption(passWord string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passWord), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// 用户密码对比
func ComparePassword(loginPwd string, userPw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPw), []byte(loginPwd))
	return err == nil
}
