package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    Category       `json:"category"`
	CategoryID  uint           `json:"category_id"`
	Level       string         `json:"level"`
}
