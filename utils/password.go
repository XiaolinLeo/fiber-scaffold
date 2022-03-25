package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPass(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("加密失败")
	}

	return string(hash)
}

func VerifyPassword(hashPass, loginPass string) bool {
	fmt.Println(EncryptPass(loginPass))
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(loginPass))
	if err != nil {
		return false
	}
	return true
}
