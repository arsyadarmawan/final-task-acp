package products

import (
	_productDomain "acp14/business/products"
	"acp14/helpers"
	"context"
	"time"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(database *gorm.DB) _productDomain.Repository {
	return &ProductRepository{
		db: database,
	}
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, data _productDomain.Domain) (int, error) {
	product := Product{
		Name:       data.Name,
		Price:      data.Price,
		CategoryId: data.CategoryId,
	}
	// fmt.Println(data)
	result := repo.db.Create(&product)
	return int(result.RowsAffected), result.Error
}

func (repo *ProductRepository) GetProducts(ctx context.Context) ([]_productDomain.Domain, error) {
	var products []Product
	result := repo.db.Find(&products)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_productDomain.Domain{}, helpers.ErrNotFound
		} else {
			return []_productDomain.Domain{}, helpers.ErrDbServer
		}
	}
	return ToListDomain(products), nil
}

func (repo *ProductRepository) DeleteProduct(ctx context.Context, id int) (int, error) {
	product := Product{}
	// fmt.Println(data)
	result := repo.db.Delete(&product, id)
	return int(result.RowsAffected), result.Error
}

func (repo *ProductRepository) UpdateProduct(ctx context.Context, id int) (int, error) {
	product := Product{}
	// fmt.Println(data)
	result := repo.db.Model(&product).Where("id", id).Update("UpdatedAt", time.Now())
	return int(result.RowsAffected), result.Error
}
