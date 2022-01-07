package databases

import (
	_cartDb "acp14/drivers/databases/carts"
	_categoryDb "acp14/drivers/databases/categories"
	_productDb "acp14/drivers/databases/products"
	_userDb "acp14/drivers/databases/users"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DBMigrate(db)
	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userDb.User{},
		&_productDb.Product{},
		&_categoryDb.Category{},
		&_cartDb.Cart{},
	)
}
