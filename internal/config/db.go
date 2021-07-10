package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnection() (*gorm.DB, error) {
	USER := viper.Get("DATABASE_USERNAME")
	PASS := viper.Get("DATABASE_PASSWORD")
	HOST := viper.Get("DATABASE_HOST")
	PORT := viper.Get("DATABASE_PORT")
	DBNAME := viper.Get("DATABASE_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err.Error())

	}
	return db, nil
}
