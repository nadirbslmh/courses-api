package users_test

import (
	"context"
	"courses-api/app/middlewares"
	"courses-api/businesses/users"
	_userMock "courses-api/businesses/users/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository _userMock.Repository
	userService    users.Usecase

	userDomain users.Domain
	ctx        context.Context
)

func TestMain(m *testing.M) {
	userService = users.NewUserUseCase(&userRepository, &middlewares.JWTConfig{})
	userDomain = users.Domain{
		Email:    "test@test.com",
		Password: "123123",
	}

	ctx = context.TODO()

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		userRepository.On("Register", ctx, &userDomain).Return(userDomain, nil).Once()

		result, err := userService.Register(ctx, &userDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register |  Invalid", func(t *testing.T) {
		userRepository.On("Register", ctx, &users.Domain{}).Return(users.Domain{}, errors.New("failed")).Once()

		result, err := userService.Register(ctx, &users.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		userRepository.On("GetByEmail", ctx, &userDomain).Return(userDomain, nil).Once()

		result, err := userService.Login(ctx, &userDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Login |  Invalid", func(t *testing.T) {
		userRepository.On("GetByEmail", ctx, &users.Domain{}).Return(users.Domain{}, errors.New("failed")).Once()

		result, err := userService.Login(ctx, &users.Domain{})

		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})
}
