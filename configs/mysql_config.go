package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/arfan21/getprint-media/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnect() (*gorm.DB, error) {
	var DBURL string

	if os.Getenv("DB_PASSWORD") == "" {
		DBURL = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	} else {
		DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	}

	var err error
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Media{})
	log.Println("MySql Connected")
	return db, nil
}