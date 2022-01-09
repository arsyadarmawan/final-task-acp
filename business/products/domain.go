package products

import (
	"context"
	"time"
)

type Domain struct {
	Id         uint
	Name       string
	Price      uint
	CategoryId uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	GetProducts(ctx context.Context) ([]Domain, error)
	CreateProduct(ctx context.Context, data Domain) (int, error)
	GetProductById(ctx context.Context, id int) (Domain, error)
	DeleteProduct(ctx context.Context, id int) (Domain, error)
	UpdateProduct(ctx context.Context, updateProduct Domain) (Domain, error)
}

type Repository interface {
	GetProducts(ctx context.Context) ([]Domain, error)
	CreateProduct(ctx context.Context, data Domain) (int, error)
	GetProductById(ctx context.Context, id int) (Domain, error)
	UpdateProduct(ctx context.Context, updateProduct Domain) (Domain, error)
	DeleteProduct(ctx context.Context, id int) (Domain, error)
}
