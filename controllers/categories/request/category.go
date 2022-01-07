package request

import (
	_categoriesDomain "acp14/business/categories"
	"time"
)

type CategoryRequest struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToDomain(category CategoryRequest) _categoriesDomain.Domain {
	return _categoriesDomain.Domain{
		Id:        category.Id,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Name:      category.Name,
	}
}
