package main

import (
	"fiber-scaffold/cmd"
	"fiber-scaffold/common/logging"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		logging.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logging.Info("Config file changed:", e.Name, e.Op)
	})
	err = cmd.Execute()

	if err != nil {
		logging.Fatal("app start err", err)
	}
}
