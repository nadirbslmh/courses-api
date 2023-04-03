package services

import (
	"courses-api/models"
	"courses-api/repositories"
)

type CategoryService struct {
	repository repositories.CategoryRepository
}

func InitCategoryService() CategoryService {
	return CategoryService{
		repository: &repositories.CategoryRepositoryImpl{},
	}
}

func (cs *CategoryService) GetAll() ([]models.Category, error) {
	return cs.repository.GetAll()
}

func (cs *CategoryService) GetByID(id string) (models.Category, error) {
	return cs.repository.GetByID(id)
}

func (cs *CategoryService) Create(categoryInput models.CategoryInput) (models.Category, error) {
	return cs.repository.Create(categoryInput)
}

func (cs *CategoryService) Update(categoryInput models.CategoryInput, id string) (models.Category, error) {
	return cs.repository.Update(categoryInput, id)
}

func (cs *CategoryService) Delete(id string) error {
	return cs.repository.Delete(id)
}
