package categories

import (
	_categoryDomain "acp14/business/categories"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        uint   `gorm:"primaryKey autoIncrement"`
	Name      string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (data *Category) ToDomain() _categoryDomain.Domain {
	return _categoryDomain.Domain{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.CreatedAt,
		Name:      data.Name,
	}
}

func ToListDomain(users []Category) []_categoryDomain.Domain {
	var result = []_categoryDomain.Domain{}
	for _, user := range users {
		result = append(result, user.ToDomain())
	}
	return result
}
