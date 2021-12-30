package users

import (
	_userDomain "acp14/business/users"
	"acp14/helpers"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) _userDomain.Repository {
	return &UserRepository{
		db: database,
	}
}

func (repo *UserRepository) GetUser(ctx context.Context) ([]_userDomain.Domain, error) {
	var users []User
	result := repo.db.Find(&users)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_userDomain.Domain{}, helpers.ErrNotFound
		} else {
			return []_userDomain.Domain{}, helpers.ErrDbServer
		}
	}
	return ToListDomain((users)), nil
}

func (repo *UserRepository) Register(ctx context.Context, data _userDomain.Domain) (int, error) {
	user := User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	result := repo.db.Create(&user)
	return int(result.RowsAffected), result.Error
}
