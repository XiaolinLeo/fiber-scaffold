package cmd

import (
	"fiber-scaffold/app/models"
	"fiber-scaffold/app/repository"
	"fiber-scaffold/db"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate db",
	Run: func(cmd *cobra.Command, args []string) {
		db.SetupMySQL()
		db.MySQL.AutoMigrate(&models.User{})
		db.MySQL.AutoMigrate(&models.Blog{})
		repository.InitUser()
	},
}
