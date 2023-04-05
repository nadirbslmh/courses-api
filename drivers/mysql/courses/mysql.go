package courses

import (
	"context"
	"courses-api/businesses/courses"

	"gorm.io/gorm"
)

type courseRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) courses.Repository {
	return &courseRepository{
		conn: conn,
	}
}

func (cr *courseRepository) GetAll(ctx context.Context) ([]courses.Domain, error) {
	var records []Course

	err := cr.conn.WithContext(ctx).Preload("Category").Find(&records).Error

	if err != nil {
		return nil, err
	}

	courses := []courses.Domain{}

	for _, course := range records {
		courses = append(courses, course.ToDomain())
	}

	return courses, nil
}

func (cr *courseRepository) GetByID(ctx context.Context, id string) (courses.Domain, error) {
	var course Course

	err := cr.conn.WithContext(ctx).Preload("Category").First(&course, "id = ?", id).Error

	if err != nil {
		return courses.Domain{}, err
	}

	return course.ToDomain(), nil
}

func (cr *courseRepository) Create(ctx context.Context, courseDomain *courses.Domain) (courses.Domain, error) {
	record := FromDomain(courseDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return courses.Domain{}, err
	}

	err := cr.conn.WithContext(ctx).Last(&record).Error

	if err != nil {
		return courses.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *courseRepository) Update(ctx context.Context, courseDomain *courses.Domain, id string) (courses.Domain, error) {
	course, err := cr.GetByID(ctx, id)

	if err != nil {
		return courses.Domain{}, err
	}

	updatedCourse := FromDomain(&course)

	updatedCourse.Title = courseDomain.Title
	updatedCourse.Description = courseDomain.Description
	updatedCourse.CategoryID = courseDomain.CategoryID
	updatedCourse.Level = courseDomain.Level

	err = cr.conn.WithContext(ctx).Save(&updatedCourse).Error

	if err != nil {
		return courses.Domain{}, err
	}

	return updatedCourse.ToDomain(), nil
}

func (cr *courseRepository) Delete(ctx context.Context, id string) error {
	course, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedCourse := FromDomain(&course)

	err = cr.conn.WithContext(ctx).Delete(&deletedCourse).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *courseRepository) Restore(ctx context.Context, id string) (courses.Domain, error) {
	var trashedCourse courses.Domain

	trashed := FromDomain(&trashedCourse)

	err := cr.conn.WithContext(ctx).Unscoped().First(&trashed, "id = ?", id).Error

	if err != nil {
		return courses.Domain{}, err
	}

	trashed.DeletedAt = gorm.DeletedAt{}

	err = cr.conn.WithContext(ctx).Unscoped().Save(&trashed).Error

	if err != nil {
		return courses.Domain{}, err
	}

	return trashed.ToDomain(), nil
}

func (cr *courseRepository) ForceDelete(ctx context.Context, id string) error {
	course, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedCourse := FromDomain(&course)

	err = cr.conn.WithContext(ctx).Unscoped().Delete(&deletedCourse).Error

	if err != nil {
		return err
	}

	return nil
}
