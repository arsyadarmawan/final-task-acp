package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Password  string
}

type Usecase interface {
	Login(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context) ([]Domain, error)
	Register(ctx context.Context, data Domain) (int, error)
}

type Repository interface {
	GetUser(ctx context.Context) ([]Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Register(ctx context.Context, data Domain) (int, error)
}
