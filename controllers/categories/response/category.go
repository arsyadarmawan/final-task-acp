package response

import (
	_categoriesDomain "acp14/business/categories"
	"time"
)

type ProductResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

func FromDomain(product _categoriesDomain.Domain) ProductResponse {
	return ProductResponse{
		Id:        product.Id,
		Name:      product.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ToListFromDomain(products []_categoriesDomain.Domain) []ProductResponse {
	var result = []ProductResponse{}
	for _, product := range products {
		result = append(result, FromDomain(product))
	}
	return result
}
