package db

import (
	"fiber-scaffold/common/logging"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MySQL *gorm.DB
var err error

func SetupMySQL() {
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_DB_NAME")
	dbHost := os.Getenv("MYSQL_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbHost, dbName)
	var config gorm.Config
	if os.Getenv("ENABLE_GORM_LOGGER") == "False" {
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
