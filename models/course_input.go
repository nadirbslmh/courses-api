package models

import "github.com/go-playground/validator/v10"

type CourseInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Category    string `json:"category" validate:"required"`
	Level       string `json:"level" validate:"required"`
}

func (c *CourseInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)

	return err
}
