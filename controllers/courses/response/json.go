package response

import (
	"courses-api/businesses/courses"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	CategoryName string         `json:"category_name"`
	CategoryID   uint           `json:"category_id"`
	Level        string         `json:"level"`
}

func FromDomain(domain courses.Domain) Course {
	return Course{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		Title:        domain.Title,
		Description:  domain.Description,
		CategoryID:   domain.CategoryID,
		CategoryName: domain.CategoryName,
		Level:        domain.Level,
	}
}
