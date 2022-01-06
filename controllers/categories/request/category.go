package request

import (
	_categoriesDomain "acp14/business/categories"
	"time"
)

type ProductRequest struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToDomain(product ProductRequest) _categoriesDomain.Domain {
	return _categoriesDomain.Domain{
		Id:        product.Id,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Name:      product.Name,
	}
}
