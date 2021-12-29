package users

import (
	_userDomain "acp14/business/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primaryKey autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	Email     string `gorm:"not null;unique"`
	Password  string
}

func (user *User) ToDomain() _userDomain.Domain {
	return _userDomain.Domain{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func ToListDomain(users []User) []_userDomain.Domain {
	var result = []_userDomain.Domain{}
	for _, user := range users {
		result = append(result, user.ToDomain())
	}
	return result
}
