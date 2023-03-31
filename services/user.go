package services

import (
	"courses-api/middlewares"
	"courses-api/models"
	"courses-api/repositories"
)

type UserService struct {
	repository repositories.UserRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitUserService(jwtAuth *middlewares.JWTConfig) UserService {
	return UserService{
		repository: &repositories.UserRepositoryImpl{},
		jwtAuth:    jwtAuth,
	}
}

func (us *UserService) Register(userInput models.UserInput) (models.User, error) {
	return us.repository.Register(userInput)
}

func (us *UserService) Login(userInput models.UserInput) (string, error) {
	user, err := us.repository.GetByEmail(userInput)

	if err != nil {
		return "", err
	}

	token, err := us.jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
