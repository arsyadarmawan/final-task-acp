package response

import (
	_productsDomain "acp14/business/products"
	"time"
)

type ProductResponse struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name"`
	Price      uint      `json:"price,omitempty"`
	CategoryId uint      `json:"category_id,omitempty"`
}

func FromDomain(product _productsDomain.Domain) ProductResponse {
	return ProductResponse{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func ToListFromDomain(products []_productsDomain.Domain) []ProductResponse {
	var result = []ProductResponse{}
	for _, product := range products {
		result = append(result, FromDomain(product))
	}
	return result
}
