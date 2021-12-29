package configs

import (
	"acp14/models/categories"
	"acp14/models/products"
	"acp14/models/users"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/acp14?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database no connected")
	}

	fmt.Println("succes ")

	// check migrate
	DB.AutoMigrate(&users.User{}, &products.Product{}, &categories.Category{})
}
