package products

import (
	_productDomain "acp14/business/products"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id         uint   `gorm:"primaryKey autoIncrement"`
	Name       string `gorm:"not null;unique"`
	Price      uint
	CategoryId uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (product *Product) ToDomain() _productDomain.Domain {
	return _productDomain.Domain{
		Id:         product.Id,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.CreatedAt,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
	}
}

func ToListDomain(users []Product) []_productDomain.Domain {
	var result = []_productDomain.Domain{}
	for _, user := range users {
		result = append(result, user.ToDomain())
	}
	return result
}
