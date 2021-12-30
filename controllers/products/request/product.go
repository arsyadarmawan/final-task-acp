package request

import (
	_productsDomain "acp14/business/products"
	"time"
)

type ProductRequest struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	Price      uint      `json:"price"`
	CategoryId uint      `json:"category_id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func ToDomain(product ProductRequest) _productsDomain.Domain {
	return _productsDomain.Domain{
		Id:         product.Id,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
		Name:       product.Name,
		Price:      uint(product.Price),
		CategoryId: uint(product.CategoryId),
	}
}
