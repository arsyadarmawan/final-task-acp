package response

import (
	_cartsDomain "acp14/business/carts"
	"time"
)

type CartResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Price     uint      `json:"price,omitempty"`
	Total     uint      `json:"total,omitempty"`
	UserId    uint      `json:"user_id,omitempty"`
	ProductId uint      `json:"product_id,omitempty"`
}

func FromDomain(data _cartsDomain.Domain) CartResponse {
	return CartResponse{
		Id:        data.Id,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		ProductId: data.ProductId,
		UserId:    data.UserId,
		Total:     data.Total,
		Price:     data.Price,
	}
}

func ToListFromDomain(carts []_cartsDomain.Domain) []CartResponse {
	var result = []CartResponse{}
	for _, cart := range carts {
		result = append(result, FromDomain(cart))
	}
	return result
}
