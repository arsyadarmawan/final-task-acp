package categories

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetCategories(ctx context.Context) ([]Domain, error)
	CreateCategory(ctx context.Context, data Domain) (int, error)
	DeleteCategory(ctx context.Context, id int) (int, error)
	UpdateCategory(ctx context.Context, id int) (int, error)
}

type Repository interface {
	GetCategories(ctx context.Context) ([]Domain, error)
	CreateCategory(ctx context.Context, data Domain) (int, error)
	DeleteCategory(ctx context.Context, id int) (int, error)
	UpdateCategory(ctx context.Context, id int) (int, error)
}
