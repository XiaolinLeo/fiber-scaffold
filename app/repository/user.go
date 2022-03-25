package repository

import (
	"fiber-scaffold/app/models"
	"fiber-scaffold/db"
	"fiber-scaffold/utils"
)

func CheckLoginUser(username, password string) bool {
	var user models.User
	if err := db.MySQL.Where("username=?", username).First(&user).Error; err != nil {
		return false
	}
	if user.ID > 0 {
		if utils.VerifyPassword(user.Password, password) {
			return true
		}
		return false
	}
	return false
}

func UserInfo(username string) string {
	var user models.User
	if err := db.MySQL.Where("username=?", username).First(&user).Error; err != nil {
		return ""
	}
	if user.ID > 0 {
		return user.Username
	}
	return ""
}

func InitUser() {
	db.MySQL.Create(&models.User{Username: "admin", Password: utils.EncryptPass("fiber")})
}
