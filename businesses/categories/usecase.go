package categories

import "context"

type categoryUsecase struct {
	categoryRepository Repository
}

func NewCategoryUsecase(repository Repository) Usecase {
	return &categoryUsecase{
		categoryRepository: repository,
	}
}

func (usecase *categoryUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.categoryRepository.GetAll(ctx)
}

func (usecase *categoryUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.categoryRepository.GetByID(ctx, id)
}

func (usecase *categoryUsecase) Create(ctx context.Context, categoryDomain *Domain) (Domain, error) {
	return usecase.categoryRepository.Create(ctx, categoryDomain)
}

func (usecase *categoryUsecase) Update(ctx context.Context, categoryDomain *Domain, id string) (Domain, error) {
	return usecase.categoryRepository.Update(ctx, categoryDomain, id)
}

func (usecase *categoryUsecase) Delete(ctx context.Context, id string) error {
	return usecase.categoryRepository.Delete(ctx, id)
}
