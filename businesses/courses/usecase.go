package courses

import "context"

type courseUsecase struct {
	courseRepository Repository
}

func NewCourseUsecase(repository Repository) Usecase {
	return &courseUsecase{
		courseRepository: repository,
	}
}

func (usecase *courseUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.courseRepository.GetAll(ctx)
}

func (usecase *courseUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.courseRepository.GetByID(ctx, id)
}

func (usecase *courseUsecase) Create(ctx context.Context, courseDomain *Domain) (Domain, error) {
	return usecase.courseRepository.Create(ctx, courseDomain)
}

func (usecase *courseUsecase) Update(ctx context.Context, courseDomain *Domain, id string) (Domain, error) {
	return usecase.courseRepository.Update(ctx, courseDomain, id)
}

func (usecase *courseUsecase) Delete(ctx context.Context, id string) error {
	return usecase.courseRepository.Delete(ctx, id)
}

func (usecase *courseUsecase) Restore(ctx context.Context, id string) (Domain, error) {
	return usecase.courseRepository.Restore(ctx, id)
}

func (usecase *courseUsecase) ForceDelete(ctx context.Context, id string) error {
	return usecase.courseRepository.ForceDelete(ctx, id)
}
