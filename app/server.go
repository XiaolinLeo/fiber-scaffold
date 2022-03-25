package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
)

func Create() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	return app
}

func Listen(app *fiber.App) error {
	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	return app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort))
}
