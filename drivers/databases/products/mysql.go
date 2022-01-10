package products

import (
	_productDomain "acp14/business/products"
	"acp14/helpers"
	"context"
	"fmt"

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

func (repo *ProductRepository) GetProductById(ctx context.Context, product_id int) (_productDomain.Domain, error) {
	var productResult Product
	result := repo.db.First(&productResult, product_id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _productDomain.Domain{}, helpers.ErrNotFound
		} else {
			return _productDomain.Domain{}, helpers.ErrDbServer
		}
	}

	return productResult.ToDomain(), nil
}

func (repo *ProductRepository) SearchCategoy(ctx context.Context, category_id int) ([]_productDomain.Domain, error) {
	var productResult []Product
	result := repo.db.Where("category_id = ?", category_id).First(&productResult)
	fmt.Println(result)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_productDomain.Domain{}, helpers.ErrNotFound
		} else {
			return []_productDomain.Domain{}, helpers.ErrDbServer
		}
	}

	return ToListDomain(productResult), nil
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

func (repo *ProductRepository) DeleteProduct(ctx context.Context, id int) (_productDomain.Domain, error) {
	var productResult Product
	// result := repo.db.Debug().Where("id = ?", id).Find(&product)
	repo.db.Where("id = ?", &id).Find(&productResult)

	result := repo.db.Delete(&Product{}, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _productDomain.Domain{}, helpers.ErrNotFound
		} else {
			return _productDomain.Domain{}, helpers.ErrDbServer
		}
	}
	return productResult.ToDomain(), nil
}

func (repo *ProductRepository) UpdateProduct(ctx context.Context, product _productDomain.Domain) (_productDomain.Domain, error) {
	result := Product{
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
	}

	query := repo.db.Model(&Product{}).Where("id = ?", product.Id).Updates(result).Find(&Product{})

	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return _productDomain.Domain{}, helpers.ErrNotFound
		} else {
			return _productDomain.Domain{}, helpers.ErrDbServer
		}
	}
	return result.ToDomain(), nil
}
