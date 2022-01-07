package carts

import (
	"context"
	"time"
)

type CartUsecase struct {
	contextTimeout time.Duration
	repo           Repository
}

func NewCartUsecase(repo Repository, timeout time.Duration) Usecase {
	return &CartUsecase{
		contextTimeout: timeout,
		repo:           repo,
	}
}

func (uc *CartUsecase) GetCarts(ctx context.Context) ([]Domain, error) {
	data, err := uc.repo.GetCarts(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (uc *CartUsecase) CreateCart(ctx context.Context, domain Domain) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.CreateCart(ctx, domain)
	return productRow, err
}

func (uc *CartUsecase) DeleteCart(ctx context.Context, id int) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.DeleteCart(ctx, id)
	return productRow, err
}

func (uc *CartUsecase) UpdateCart(ctx context.Context, id int) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.UpdateCart(ctx, id)
	return productRow, err
}
