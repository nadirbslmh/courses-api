package courses

import (
	"courses-api/businesses/courses"
	"courses-api/drivers/mysql/categories"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Category    categories.Category `json:"category"`
	CategoryID  uint                `json:"category_id"`
	Level       string              `json:"level"`
}

func (rec *Course) ToDomain() courses.Domain {
	return courses.Domain{
		ID:           rec.ID,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
		Title:        rec.Title,
		Description:  rec.Description,
		CategoryName: rec.Category.Name,
		CategoryID:   rec.Category.ID,
		Level:        rec.Level,
	}
}

func FromDomain(domain *courses.Domain) *Course {
	return &Course{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		Title:       domain.Title,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
		Level:       domain.Level,
	}
}
