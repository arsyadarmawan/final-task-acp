package carts

import (
	_cartDomain "acp14/business/carts"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Id        uint   `gorm:"primaryKey autoIncrement"`
	Name      string `gorm:"not null;unique"`
	Price     uint
	ProductId uint
	UserId    uint
	Total     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (data *Cart) ToDomain() _cartDomain.Domain {
	return _cartDomain.Domain{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.CreatedAt,
		Name:      data.Name,
		Price:     data.Price,
		Total:     data.Total,
		ProductId: data.ProductId,
	}
}

func ToListDomain(users []Cart) []_cartDomain.Domain {
	var result = []_cartDomain.Domain{}
	for _, user := range users {
		result = append(result, user.ToDomain())
	}
	return result
}
