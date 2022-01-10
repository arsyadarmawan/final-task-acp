package request

import (
	_cartsDomain "acp14/business/carts"
	"time"
)

type DeleteCart struct {
	Id uint `param:"id"`
}

type CartRequest struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     uint      `json:"price"`
	Total     uint      `json:"total"`
	UserId    uint      `json:"user_id"`
	ProductId uint      `json:"product_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateRequest struct {
	Id    uint `param:"id"`
	Total uint `json:"total"`
}

func ToDomain(data CartRequest) _cartsDomain.Domain {
	return _cartsDomain.Domain{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Name:      data.Name,
		UserId:    uint(data.UserId),
		Price:     uint(data.Price),
		ProductId: uint(data.ProductId),
		Total:     uint(data.Total),
	}
}
