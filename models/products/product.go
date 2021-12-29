package products

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint   `gorm:"primaryKey autoIncrement"`
	Name       string `gorm:"not null"`
	Desription string `gorm:"not null"`
	Price      int64
	Image      string
	CategoryId uint
	// Category   categories.Category
	// category      categories.Category `gorm:"foreignKey:categoryRefer"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeteledAt gorm.DeletedAt `gorm:"index"`
}
