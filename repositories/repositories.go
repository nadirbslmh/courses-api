package repositories

import "courses-api/models"

type CourseRepository interface {
	GetAll() []models.Course
	GetByID(id string) (models.Course, error)
	Create(courseInput models.Course) (models.Course, error)
	Update(courseInput models.Course, id string) (models.Course, error)
	Delete(id string) error
}
