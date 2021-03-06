package carts

import (
	_cartsDomain "acp14/business/carts"
	"acp14/helpers"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) _cartsDomain.Repository {
	return &CartRepository{
		db: database,
	}
}

func (repo *CartRepository) CreateCart(ctx context.Context, data _cartsDomain.Domain) (int, error) {
	cart := Cart{
		Name:  data.Name,
		Total: data.Total,

		Price:     data.Price,
		ProductId: data.ProductId,
	}
	fmt.Println(cart.UserId)
	result := repo.db.Create(&cart)
	return int(result.RowsAffected), result.Error
}

func (repo *CartRepository) GetCarts(ctx context.Context) ([]_cartsDomain.Domain, error) {
	var carts []Cart
	result := repo.db.Find(&carts)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_cartsDomain.Domain{}, helpers.ErrNotFound
		} else {
			return []_cartsDomain.Domain{}, helpers.ErrDbServer
		}
	}
	return ToListDomain(carts), nil
}

func (repo *CartRepository) DeleteCart(ctx context.Context, id int) (int, error) {
	cart := Cart{}
	// fmt.Println(data)
	result := repo.db.Delete(&cart, id)
	return int(result.RowsAffected), result.Error
}

func (repo *CartRepository) UpdateCart(ctx context.Context, domain _cartsDomain.Domain) (_cartsDomain.Domain, error) {
	cart := Cart{
		Total: domain.Total,
	}
	query := repo.db.Model(&Cart{}).Where("id = ?", domain.Id).Updates(cart).Find(&Cart{})

	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return _cartsDomain.Domain{}, helpers.ErrNotFound
		} else {
			return _cartsDomain.Domain{}, helpers.ErrDbServer
		}
	}
	return cart.ToDomain(), nil
}
