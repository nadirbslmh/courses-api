package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email     string
	Password  string
}

type Usecase interface {
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	Login(ctx context.Context, userDomain *Domain) (string, error)
}

type Repository interface {
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
}
