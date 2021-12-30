package products

import (
	"context"
	"time"
)

type Domain struct {
	Id         uint
	Name       string
	Price      string
	CategoryId uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	GetProducts(ctx context.Context) ([]Domain, error)
	CreateProduct(ctx context.Context, data Domain) (int, error)
	DeleteProduct(ctx context.Context, id int) (int, error)
}

type Repository interface {
	GetProducts(ctx context.Context) ([]Domain, error)
	CreateProduct(ctx context.Context, data Domain) (int, error)
	DeleteProduct(ctx context.Context, id int) (int, error)
}
