package models

import "github.com/go-playground/validator/v10"

type CategoryInput struct {
	Name string `json:"name" validate:"required"`
}

func (c *CategoryInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)

	return err
}
