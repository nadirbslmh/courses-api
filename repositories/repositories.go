package repositories

import "courses-api/models"

type CourseRepository interface {
	GetAll() ([]models.Course, error)
	GetByID(id string) (models.Course, error)
	Create(courseInput models.CourseInput) (models.Course, error)
	Update(courseInput models.CourseInput, id string) (models.Course, error)
	Delete(id string) error
	Restore(id string) (models.Course, error)
	ForceDelete(id string) error
}

type UserRepository interface {
	Register(userInput models.UserInput) (models.User, error)
	GetByEmail(userInput models.UserInput) (models.User, error)
}

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id string) (models.Category, error)
	Create(categoryInput models.CategoryInput) (models.Category, error)
	Update(categoryInput models.CategoryInput, id string) (models.Category, error)
	Delete(id string) error
}
