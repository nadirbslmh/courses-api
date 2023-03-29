package services

import (
	"courses-api/models"
	"courses-api/repositories"
)

type CourseService struct {
	repository repositories.CourseRepository
}

func InitCourseService() CourseService {
	return CourseService{
		repository: &repositories.CourseRepositoryImpl{},
	}
}

func (cs *CourseService) GetAll() ([]models.Course, error) {
	return cs.repository.GetAll()
}

func (cs *CourseService) GetByID(id string) (models.Course, error) {
	return cs.repository.GetByID(id)
}

func (cs *CourseService) Create(courseInput models.CourseInput) (models.Course, error) {
	return cs.repository.Create(courseInput)
}

func (cs *CourseService) Update(courseInput models.CourseInput, id string) (models.Course, error) {
	return cs.repository.Update(courseInput, id)
}

func (cs *CourseService) Delete(id string) error {
	return cs.repository.Delete(id)
}

func (cs *CourseService) Restore(id string) (models.Course, error) {
	return cs.repository.Restore(id)
}

func (cs *CourseService) ForceDelete(id string) error {
	return cs.repository.ForceDelete(id)
}
