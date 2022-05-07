package utils

import (
	"strings"

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

// 用户手机号码加密
func PhoneNoCryption(phoneNo string) string {
	var cryptionPhone []string
	if phoneNo == "" {
		return ""
	} else {
		phoneNo := phoneNo[:len(phoneNo)-4]
		cryptionPhone = append(cryptionPhone, phoneNo)
		cryptionPhone = append(cryptionPhone, "****")
		return strings.Join(cryptionPhone, "")
	}
}
