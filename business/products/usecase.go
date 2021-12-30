package products

import (
	"context"
	"time"
)

type ProductUsecase struct {
	contextTimeout time.Duration
	repo           Repository
}

func NewProductUsecase(repo Repository, timeout time.Duration) Usecase {
	return &ProductUsecase{
		contextTimeout: timeout,
		repo:           repo,
	}
}

func (uc *ProductUsecase) GetProducts(ctx context.Context) ([]Domain, error) {
	data, err := uc.repo.GetProducts(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, domain Domain) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.CreateProduct(ctx, domain)
	return productRow, err
}

func (uc *ProductUsecase) DeleteProduct(ctx context.Context, id int) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.DeleteProduct(ctx, id)
	return productRow, err
}

func (uc *ProductUsecase) UpdateProduct(ctx context.Context, id int) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.UpdateProduct(ctx, id)
	return productRow, err
}
