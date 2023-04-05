package users

import (
	"courses-api/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Email:     rec.Email,
		Password:  rec.Password,
	}
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Email:     domain.Email,
		Password:  domain.Password,
	}
}
