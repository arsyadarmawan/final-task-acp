package categories

import (
	"context"
	"time"
)

type CategoryUsecase struct {
	contextTimeout time.Duration
	repo           Repository
}

func NewCategoryUsecase(repo Repository, timeout time.Duration) Usecase {
	return &CategoryUsecase{
		contextTimeout: timeout,
		repo:           repo,
	}
}

func (uc *CategoryUsecase) GetCategories(ctx context.Context) ([]Domain, error) {
	data, err := uc.repo.GetCategories(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (uc *CategoryUsecase) CreateCategory(ctx context.Context, domain Domain) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.CreateCategory(ctx, domain)
	return productRow, err
}

func (uc *CategoryUsecase) DeleteCategory(ctx context.Context, id int) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.DeleteCategory(ctx, id)
	return productRow, err
}

func (uc *CategoryUsecase) UpdateCategory(ctx context.Context, id int) (int, error) {
	var productRow int
	var err error
	productRow, err = uc.repo.UpdateCategory(ctx, id)
	return productRow, err
}
