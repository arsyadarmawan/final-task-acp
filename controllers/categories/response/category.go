package response

import (
	_categoriesDomain "acp14/business/categories"
	"time"
)

type CategoryResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

func FromDomain(product _categoriesDomain.Domain) CategoryResponse {
	return CategoryResponse{
		Id:        product.Id,
		Name:      product.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ToListFromDomain(products []_categoriesDomain.Domain) []CategoryResponse {
	var result = []CategoryResponse{}
	for _, product := range products {
		result = append(result, FromDomain(product))
	}
	return result
}
