package users

import (
	"context"
	"errors"
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

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	var err error
	var user Domain
	user, err = uc.repository.GetByEmail(ctx, email)

	if err != nil {
		return "user not found", err
	}

	if user.Password != password {
		return "password not match to our record", errors.New("password not match to our record")
	} else {
		return "match", nil
	}
}
