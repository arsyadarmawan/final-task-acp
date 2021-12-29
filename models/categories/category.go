package categories

import (
	"acp14/models/products"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Id        uint `gorm:"primaryKey autoIncrement"`
	name      string
	products  []products.Product
	CreatedAt time.Time
	UpdatedAt time.Time
	DeteledAt gorm.DeletedAt `gorm:"index"`
}
