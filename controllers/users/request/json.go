package request

import (
	"courses-api/businesses/users"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
