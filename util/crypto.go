package util

import "golang.org/x/crypto/bcrypt"

func Encrypt(origin string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(origin), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

//比较加密后的密码和未加密的密码，如果一致返回true，如果不一致返回false
func ValidatePassword(crypted, passwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(crypted), []byte(passwd))
	if err != nil {
		return false
	}
	return true
}
