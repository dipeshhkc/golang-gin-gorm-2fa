package database

import (
	"fmt"
	"golang-gin-gorm-2fa/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//DBConnection -> return db instance
func DBConnection() *gorm.DB {
	USER := "root"
	PASS := "root"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "2fa"

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
		panic(err)
	}
	db.AutoMigrate(&model.TwoFactor{})
	return db

}
