package users

import (
	"context"
	"time"
)

type UserUsecase struct {
	contextTimeout time.Duration
	repository     Repository
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *UserUsecase) GetUser(ctx context.Context) ([]Domain, error) {
	users, err := uc.repository.GetUser(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return users, nil
}

func (uc *UserUsecase) Register(ctx context.Context, domain Domain) (int, error) {
	val, err := uc.repository.Register(ctx, domain)
	return val, err
}
