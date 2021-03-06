package app

import (
	b "fiber-scaffold/app/handler/blog"
	"fiber-scaffold/app/handler/user"
	"fiber-scaffold/app/middleware"

	"github.com/gofiber/fiber/v2"
)

// Embed a directory
////go:embed web/*
//var embedDirStatic embed.FS

func SetupRouters(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	admin := app.Group("/user")
	admin.Post("/login", user.UserLogin)
	admin.Post("/info", user.UserInfo)
	blog := app.Group("/blog", middleware.Auth())
	blog.Post("/add", b.AddBlog)
}
