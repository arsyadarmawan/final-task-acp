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

func (uc *ProductUsecase) GetProductById(ctx context.Context, id int) (Domain, error) {
	Product, err := uc.repo.GetProductById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return Product, nil
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, domain Domain) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.CreateProduct(ctx, domain)
	return productRow, err
}

func (uc *ProductUsecase) DeleteProduct(ctx context.Context, id int) (Domain, error) {
	product, err := uc.repo.DeleteProduct(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return product, nil
}

func (uc *ProductUsecase) UpdateProduct(ctx context.Context, product Domain) (Domain, error) {
	result, err := uc.repo.UpdateProduct(ctx, product)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (uc *ProductUsecase) SearchCategoy(ctx context.Context, id int) ([]Domain, error) {
	result, err := uc.repo.SearchCategoy(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}
