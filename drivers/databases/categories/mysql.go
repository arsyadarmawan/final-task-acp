package categories

import (
	_categoriesDomain "acp14/business/categories"
	"acp14/helpers"
	"context"
	"time"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(database *gorm.DB) _categoriesDomain.Repository {
	return &CategoryRepository{
		db: database,
	}
}

func (repo *CategoryRepository) CreateCategory(ctx context.Context, data _categoriesDomain.Domain) (int, error) {
	category := Category{
		Name: data.Name,
	}
	// fmt.Println(data)
	result := repo.db.Create(&category)
	return int(result.RowsAffected), result.Error
}

func (repo *CategoryRepository) GetCategories(ctx context.Context) ([]_categoriesDomain.Domain, error) {
	var categories []Category
	result := repo.db.Find(&categories)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_categoriesDomain.Domain{}, helpers.ErrNotFound
		} else {
			return []_categoriesDomain.Domain{}, helpers.ErrDbServer
		}
	}
	return ToListDomain(categories), nil
}

func (repo *CategoryRepository) DeleteCategory(ctx context.Context, id int) (int, error) {
	category := Category{}
	// fmt.Println(data)
	result := repo.db.Delete(&category, id)
	return int(result.RowsAffected), result.Error
}

func (repo *CategoryRepository) UpdateCategory(ctx context.Context, id int) (int, error) {
	category := Category{}
	// fmt.Println(data)
	result := repo.db.Model(&category).Where("id", id).Update("UpdatedAt", time.Now())
	return int(result.RowsAffected), result.Error
}
