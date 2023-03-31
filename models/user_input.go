package models

import "github.com/go-playground/validator/v10"

type UserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *UserInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}
