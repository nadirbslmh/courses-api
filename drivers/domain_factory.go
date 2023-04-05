package drivers

import (
	categoryDomain "courses-api/businesses/categories"
	categoryDB "courses-api/drivers/mysql/categories"

	courseDomain "courses-api/businesses/courses"
	courseDB "courses-api/drivers/mysql/courses"

	userDomain "courses-api/businesses/users"
	userDB "courses-api/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewCourseRepository(conn *gorm.DB) courseDomain.Repository {
	return courseDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
