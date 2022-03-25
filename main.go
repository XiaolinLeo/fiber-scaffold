package main

import (
	"fiber-scaffold/cmd"
	"fiber-scaffold/common/logging"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logging.Info("Error loading .env file")
	}

	//app := fiber.New(fiber.Config{
	//	DisableStartupMessage: false,
	//})

	err = cmd.Execute()

	if err != nil {
		logging.Fatal("start err", err)
	}
}
