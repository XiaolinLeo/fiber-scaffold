package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `json:"title" gorm:"title"`
	Content string `json:"content" gorm:"type:text"`
}
