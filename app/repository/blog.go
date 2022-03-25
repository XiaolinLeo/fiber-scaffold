package repository

import (
	"fiber-scaffold/app/models"
	"fiber-scaffold/db"
)

func AddBlog(blog *models.Blog) error {
	if err := db.MySQL.Create(&blog).Error; err != nil {
		return err
	}
	return nil
}
