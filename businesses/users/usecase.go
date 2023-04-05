package users

import (
	"context"
	"courses-api/app/middlewares"
)

type userUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.JWTConfig
}

func NewUserUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &userUsecase{
		userRepository: repository,
		jwtAuth:        jwtAuth,
	}
}

func (usecase *userUsecase) Register(ctx context.Context, userDomain *Domain) (Domain, error) {
	return usecase.userRepository.Register(ctx, userDomain)
}

func (usecase *userUsecase) Login(ctx context.Context, userDomain *Domain) (string, error) {
	user, err := usecase.userRepository.GetByEmail(ctx, userDomain)

	if err != nil {
		return "", err
	}

	token, err := usecase.jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
