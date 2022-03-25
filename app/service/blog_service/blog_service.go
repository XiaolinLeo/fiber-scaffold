package blog_service

import (
	"fiber-scaffold/app/models"
	"fiber-scaffold/app/repository"
	"time"
)

type Blog struct {
	ID      int64
	Title   string
	Content string
	AddTime time.Time
}

func (b Blog) Add() error {
	blog := &models.Blog{
		Title:   b.Title,
		Content: b.Content,
	}
	err := repository.AddBlog(blog)
	if err != nil {
		return err
	}

	return nil
}
