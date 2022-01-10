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

type DeleteProduct struct {
	Id uint `param:"id"`
}

type UpdateProduct struct {
	Id         uint   `param:"id"`
	Name       string `json:"name"`
	CategoryId uint   `json:"category_id"`
	Price      int    `json:"price"`
}

type ProductId struct {
	Id uint `param:"product_id"`
}

type CategoryId struct {
	Id uint `param:"category_id"`
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
