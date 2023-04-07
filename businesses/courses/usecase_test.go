package courses_test

import (
	"context"
	"courses-api/businesses/categories"
	"courses-api/businesses/courses"
	_courseMock "courses-api/businesses/courses/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	courseRepository _courseMock.Repository
	courseService    courses.Usecase

	courseDomain courses.Domain
	ctx          context.Context
)

func TestMain(m *testing.M) {
	courseService = courses.NewCourseUsecase(&courseRepository)

	categoryDomain := categories.Domain{
		Name: "test",
	}

	courseDomain = courses.Domain{
		Title:       "test",
		Description: "test",
		Level:       "all levels",
		CategoryID:  categoryDomain.ID,
	}

	ctx = context.TODO()

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		courseRepository.On("GetAll", ctx).Return([]courses.Domain{courseDomain}, nil).Once()

		result, err := courseService.GetAll(ctx)

		assert.Equal(t, 1, len(result))
		assert.Nil(t, err)
	})

	t.Run("Get All |  Invalid", func(t *testing.T) {
		courseRepository.On("GetAll", ctx).Return([]courses.Domain{}, nil).Once()

		result, err := courseService.GetAll(ctx)

		assert.Equal(t, 0, len(result))
		assert.Nil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		courseRepository.On("GetByID", ctx, "1").Return(courseDomain, nil).Once()

		result, err := courseService.GetByID(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get By ID |  Invalid", func(t *testing.T) {
		courseRepository.On("GetByID", ctx, "-1").Return(courses.Domain{}, errors.New("failed")).Once()

		result, err := courseService.GetByID(ctx, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		courseRepository.On("Create", ctx, &courseDomain).Return(courseDomain, nil).Once()

		result, err := courseService.Create(ctx, &courseDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create |  Invalid", func(t *testing.T) {
		courseRepository.On("Create", ctx, &courses.Domain{}).Return(courses.Domain{}, errors.New("failed")).Once()

		result, err := courseService.Create(ctx, &courses.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		courseRepository.On("Update", ctx, &courseDomain, "1").Return(courseDomain, nil).Once()

		result, err := courseService.Update(ctx, &courseDomain, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update |  Invalid", func(t *testing.T) {
		courseRepository.On("Update", ctx, &courses.Domain{}, "1").Return(courses.Domain{}, errors.New("failed")).Once()

		result, err := courseService.Update(ctx, &courses.Domain{}, "1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		courseRepository.On("Delete", ctx, "1").Return(nil).Once()

		err := courseService.Delete(ctx, "1")

		assert.Nil(t, err)
	})

	t.Run("Delete |  Invalid", func(t *testing.T) {
		courseRepository.On("Delete", ctx, "-1").Return(errors.New("failed")).Once()

		err := courseService.Delete(ctx, "-1")

		assert.NotNil(t, err)
	})
}

func TestRestore(t *testing.T) {
	t.Run("Restore | Valid", func(t *testing.T) {
		courseRepository.On("Restore", ctx, "1").Return(courseDomain, nil).Once()

		result, err := courseService.Restore(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Restore |  Invalid", func(t *testing.T) {
		courseRepository.On("Restore", ctx, "-1").Return(courses.Domain{}, errors.New("failed")).Once()

		result, err := courseService.Restore(ctx, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestForceDelete(t *testing.T) {
	t.Run("Force Delete | Valid", func(t *testing.T) {
		courseRepository.On("ForceDelete", ctx, "1").Return(nil).Once()

		err := courseService.ForceDelete(ctx, "1")

		assert.Nil(t, err)
	})

	t.Run("Force Delete |  Invalid", func(t *testing.T) {
		courseRepository.On("ForceDelete", ctx, "-1").Return(errors.New("failed")).Once()

		err := courseService.ForceDelete(ctx, "-1")

		assert.NotNil(t, err)
	})
}
