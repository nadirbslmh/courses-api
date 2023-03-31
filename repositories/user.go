package repositories

import (
	"courses-api/database"
	"courses-api/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
}

func InitUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) Register(userInput models.UserInput) (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Email:    userInput.Email,
		Password: string(password),
	}

	result := database.DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	err = result.Last(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetByEmail(userInput models.UserInput) (models.User, error) {
	var user models.User

	err := database.DB.First(&user, "email = ?", userInput.Email).Error

	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
