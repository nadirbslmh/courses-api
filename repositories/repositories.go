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
