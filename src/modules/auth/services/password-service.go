package auth_services

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(plainText string) string {
    bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)
	if err != nil{
		panic(err);
	}
    return string(bytes)
}

func IsPasswordValid(cipher string, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(cipher), []byte(password))
    return err == nil
}
