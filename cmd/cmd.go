package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(startAppCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
