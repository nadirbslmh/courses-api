package repositories

import (
	"courses-api/models"
	"errors"
)

var database []models.Course

type CourseRepositoryImpl struct{}

func InitCourseRepository() CourseRepository {
	return &CourseRepositoryImpl{}
}

func (cr *CourseRepositoryImpl) GetAll() []models.Course {
	return database
}

func (cr *CourseRepositoryImpl) GetByID(id string) (models.Course, error) {
	var foundCourse models.Course
	var isFound bool = false

	for _, course := range database {
		if course.ID == id {
			foundCourse = course
			isFound = true
		}
	}

	if isFound == false {
		return models.Course{}, errors.New("course not found")
	}

	return foundCourse, nil
}

func (cr *CourseRepositoryImpl) Create(courseInput models.Course) (models.Course, error) {
	// validate
	if courseInput.Title == "" || courseInput.Description == "" {
		return models.Course{}, errors.New("invalid fields")
	}

	var createdCourse models.Course = models.Course{
		ID:          courseInput.ID,
		Title:       courseInput.Title,
		Description: courseInput.Description,
		Category:    courseInput.Category,
		Level:       courseInput.Level,
	}

	database = append(database, createdCourse)

	return createdCourse, nil
}

func (cr *CourseRepositoryImpl) Update(courseInput models.Course, id string) (models.Course, error) {
	if courseInput.Title == "" || courseInput.Description == "" {
		return models.Course{}, errors.New("invalid fields")
	}

	var updatedCourse models.Course

	for _, course := range database {
		if course.ID == id {
			course.Title = courseInput.Title
			course.Description = courseInput.Description
			course.Category = courseInput.Category
			course.Level = courseInput.Level
			updatedCourse = course
		}
	}

	return updatedCourse, nil
}

func (cr *CourseRepositoryImpl) Delete(id string) error {
	var isFound bool = false

	for idx, course := range database {
		if course.ID == id {
			database = append(database, database[:idx]...)
			isFound = true
		}
	}

	if isFound == false {
		return errors.New("course not found")
	}

	return nil
}
