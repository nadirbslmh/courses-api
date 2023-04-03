package repositories

import (
	"courses-api/database"
	"courses-api/models"

	"gorm.io/gorm"
)

type CourseRepositoryImpl struct{}

func InitCourseRepository() CourseRepository {
	return &CourseRepositoryImpl{}
}

func (cr *CourseRepositoryImpl) GetAll() ([]models.Course, error) {
	var courses []models.Course

	err := database.DB.Preload("Category").Find(&courses).Error

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (cr *CourseRepositoryImpl) GetByID(id string) (models.Course, error) {
	var course models.Course

	err := database.DB.Preload("Category").First(&course, "id = ?", id).Error

	if err != nil {
		return models.Course{}, err
	}

	return course, nil
}

func (cr *CourseRepositoryImpl) Create(courseInput models.CourseInput) (models.Course, error) {
	var createdCourse models.Course = models.Course{
		Title:       courseInput.Title,
		Description: courseInput.Description,
		CategoryID:  courseInput.CategoryID,
		Level:       courseInput.Level,
	}

	result := database.DB.Create(&createdCourse)

	if err := result.Error; err != nil {
		return models.Course{}, err
	}

	err := database.DB.Last(&createdCourse).Error

	if err != nil {
		return models.Course{}, err
	}

	return createdCourse, nil
}

func (cr *CourseRepositoryImpl) Update(courseInput models.CourseInput, id string) (models.Course, error) {
	course, err := cr.GetByID(id)

	if err != nil {
		return models.Course{}, err
	}

	course.Title = courseInput.Title
	course.Description = courseInput.Description
	course.CategoryID = courseInput.CategoryID
	course.Level = courseInput.Level

	err = database.DB.Save(&course).Error

	if err != nil {
		return models.Course{}, err
	}

	return course, nil
}

func (cr *CourseRepositoryImpl) Delete(id string) error {
	course, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	err = database.DB.Delete(&course).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *CourseRepositoryImpl) Restore(id string) (models.Course, error) {
	var trashedCourse models.Course

	err := database.DB.Unscoped().First(&trashedCourse, "id = ?", id).Error

	if err != nil {
		return models.Course{}, err
	}

	trashedCourse.DeletedAt = gorm.DeletedAt{}

	err = database.DB.Unscoped().Save(&trashedCourse).Error

	if err != nil {
		return models.Course{}, err
	}

	return trashedCourse, nil
}

func (cr *CourseRepositoryImpl) ForceDelete(id string) error {
	course, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	err = database.DB.Unscoped().Delete(&course).Error

	if err != nil {
		return err
	}

	return nil
}
