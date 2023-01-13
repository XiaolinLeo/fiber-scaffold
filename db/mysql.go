package db

import (
	"fiber-scaffold/common/logging"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var MySQL *gorm.DB
var err error

func SetupMySQL() {
	username := viper.GetString("MYSQL.USER")
	password := viper.GetString("MYSQL.PASS")
	dbName := viper.GetString("MYSQL.DB_NAME")
	dbHost := viper.GetString("MYSQL.HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbName)
	var config gorm.Config
	if viper.GetString("MYSQL.ENABLE_GORM_LOGGER") == "False" {
		config = gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
	} else {
		config = gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	MySQL, err = gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		logging.Fatal(err)
	}
	log.Print("connect success")
}
