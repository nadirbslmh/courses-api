package database

import (
	"courses-api/models"
	"courses-api/utils"
	"errors"
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
	err := DB.AutoMigrate(&models.Course{}, &models.User{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}

func SeedCategory() (models.Category, error) {
	var category models.Category = models.Category{
		Name: "test",
	}

	result := DB.Create(&category)

	if err := result.Error; err != nil {
		return models.Category{}, err
	}

	if err := result.Last(&category).Error; err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func SeedCourse() (models.Course, error) {
	category, err := SeedCategory()

	if err != nil {
		return models.Course{}, err
	}

	var course models.Course = models.Course{
		Title:       "test",
		Description: "test",
		CategoryID:  category.ID,
		Level:       "test",
	}

	result := DB.Create(&course)

	if err := result.Error; err != nil {
		return models.Course{}, err
	}

	if err := result.Last(&course).Error; err != nil {
		return models.Course{}, err
	}

	return course, nil
}

func CleanSeeders() error {
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	catErr := DB.Exec("DELETE FROM categories").Error
	courseErr := DB.Exec("DELETE FROM courses").Error

	var isFailed bool = catErr != nil || courseErr != nil

	if isFailed {
		return errors.New("cleaning failed")
	}

	log.Println("seeders are cleaned up successfully")

	return nil
}
