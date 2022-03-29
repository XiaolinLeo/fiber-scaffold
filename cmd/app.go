package cmd

import (
	"fiber-scaffold/app"
	"fiber-scaffold/db"
	"fiber-scaffold/utils"
	"github.com/spf13/cobra"
	"log"
)

var startAppCmd = &cobra.Command{
	Use:   "start",
	Short: "start app server",
	Run: func(cmd *cobra.Command, args []string) {
		db.SetupMySQL()
		db.SetupRedis()
		utils.SetupCron()
		server := app.Create()
		app.SetupRouters(server)

		if err := app.Listen(server); err != nil {
			log.Panic(err)
		}
	},
}
