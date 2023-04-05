package request

import (
	"courses-api/businesses/courses"

	"github.com/go-playground/validator/v10"
)

type Course struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	Level       string `json:"level" validate:"required"`
}

func (req *Course) ToDomain() *courses.Domain {
	return &courses.Domain{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Level:       req.Level,
	}
}

func (req *Course) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
