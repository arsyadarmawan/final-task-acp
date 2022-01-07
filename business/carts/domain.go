package carts

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	Name      string
	Price     uint
	ProductId uint
	Total     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetCarts(ctx context.Context) ([]Domain, error)
	CreateCart(ctx context.Context, data Domain) (int, error)
	DeleteCart(ctx context.Context, id int) (int, error)
	UpdateCart(ctx context.Context, id int) (int, error)
}

type Repository interface {
	GetCarts(ctx context.Context) ([]Domain, error)
	CreateCart(ctx context.Context, data Domain) (int, error)
	DeleteCart(ctx context.Context, id int) (int, error)
	UpdateCart(ctx context.Context, id int) (int, error)
}
