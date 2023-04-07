package categories_test

import (
	"context"
	"courses-api/businesses/categories"
	_categoryMock "courses-api/businesses/categories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoryRepository _categoryMock.Repository
	categoryService    categories.Usecase

	categoryDomain categories.Domain
	ctx            context.Context
)

func TestMain(m *testing.M) {
	categoryService = categories.NewCategoryUsecase(&categoryRepository)
	categoryDomain = categories.Domain{
		Name: "test",
	}

	ctx = context.TODO()

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll", ctx).Return([]categories.Domain{categoryDomain}, nil).Once()

		result, err := categoryService.GetAll(ctx)

		assert.Equal(t, 1, len(result))
		assert.Nil(t, err)
	})

	t.Run("Get All |  Invalid", func(t *testing.T) {
		categoryRepository.On("GetAll", ctx).Return([]categories.Domain{}, nil).Once()

		result, err := categoryService.GetAll(ctx)

		assert.Equal(t, 0, len(result))
		assert.Nil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		categoryRepository.On("GetByID", ctx, "1").Return(categoryDomain, nil).Once()

		result, err := categoryService.GetByID(ctx, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Get By ID |  Invalid", func(t *testing.T) {
		categoryRepository.On("GetByID", ctx, "-1").Return(categories.Domain{}, errors.New("failed")).Once()

		result, err := categoryService.GetByID(ctx, "-1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		categoryRepository.On("Create", ctx, &categoryDomain).Return(categoryDomain, nil).Once()

		result, err := categoryService.Create(ctx, &categoryDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create |  Invalid", func(t *testing.T) {
		categoryRepository.On("Create", ctx, &categories.Domain{}).Return(categories.Domain{}, errors.New("failed")).Once()

		result, err := categoryService.Create(ctx, &categories.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		categoryRepository.On("Update", ctx, &categoryDomain, "1").Return(categoryDomain, nil).Once()

		result, err := categoryService.Update(ctx, &categoryDomain, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update |  Invalid", func(t *testing.T) {
		categoryRepository.On("Update", ctx, &categories.Domain{}, "1").Return(categories.Domain{}, errors.New("failed")).Once()

		result, err := categoryService.Update(ctx, &categories.Domain{}, "1")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		categoryRepository.On("Delete", ctx, "1").Return(nil).Once()

		err := categoryService.Delete(ctx, "1")

		assert.Nil(t, err)
	})

	t.Run("Delete |  Invalid", func(t *testing.T) {
		categoryRepository.On("Delete", ctx, "-1").Return(errors.New("failed")).Once()

		err := categoryService.Delete(ctx, "-1")

		assert.NotNil(t, err)
	})
}
