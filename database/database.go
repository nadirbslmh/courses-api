package database

import (
	"courses-api/models"
	"courses-api/utils"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = utils.GetConfig("DB_USERNAME")
	DB_PASSWORD string = utils.GetConfig("DB_PASSWORD")
	DB_NAME     string = utils.GetConfig("DB_NAME")
	DB_HOST     string = utils.GetConfig("DB_HOST")
	DB_PORT     string = utils.GetConfig("DB_PORT")
)

// connect to the database
func InitDatabase() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to the database")
}

// perform migration
func Migrate() {
	err := DB.AutoMigrate(&models.Course{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}
