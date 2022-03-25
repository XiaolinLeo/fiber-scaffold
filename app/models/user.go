package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"user_name,unique"`
	Password string `json:"password"`
}
