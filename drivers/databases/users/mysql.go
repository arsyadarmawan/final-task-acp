package users

import (
	_userDomain "acp14/business/users"
	"acp14/helpers"
	"context"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
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
		Password: HashPassword(data.Password),
	}
	result := repo.db.Create(&user)
	return int(result.RowsAffected), result.Error
}

func (repo *UserRepository) GetByEmail(ctx context.Context, email string) (_userDomain.Domain, error) {
	var user User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		// fmt.Println(result.Error)
		return _userDomain.Domain{}, result.Error
		// if result.Error == gorm.ErrRecordNotFound {

		// } else {
		// 	return _userDomain.Domain{}, helpers.ErrDbServer
		// }
	}
	return _userDomain.Domain{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
