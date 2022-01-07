package main

import (
	_cartUsecase "acp14/business/carts"
	_categoryUsecase "acp14/business/categories"
	_productUsecase "acp14/business/products"
	_userUsecase "acp14/business/users"
	_cartController "acp14/controllers/carts"
	_categoryController "acp14/controllers/categories"
	_productController "acp14/controllers/products"
	_userController "acp14/controllers/users"
	_cartRepo "acp14/drivers/databases/carts"
	_categoryRepo "acp14/drivers/databases/categories"
	_productRepo "acp14/drivers/databases/products"
	_userRepo "acp14/drivers/databases/users"
	"fmt"

	_dbDriver "acp14/drivers/databases"

	_route "acp14/apps/routes"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("apps/configs/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}
func main() {
	// Start Echo
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	fmt.Println(configDB)

	db := configDB.ConnectDB()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepo := _userRepo.NewUserRepository(db)
	userUseCase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	userController := _userController.NewUserController(userUseCase)

	// product repo
	productRepo := _productRepo.NewProductRepository(db)
	productUseCase := _productUsecase.NewProductUsecase(productRepo, timeoutContext)
	productController := _productController.NewProductController(productUseCase)

	// category repo
	categoryrepo := _categoryRepo.NewCategoryRepository(db)
	categoryUseCase := _categoryUsecase.NewCategoryUsecase(categoryrepo, timeoutContext)
	categoryController := _categoryController.NewCategoryController(categoryUseCase)

	// cart Repor
	cartRepo := _cartRepo.NewCategoryRepository(db)
	cartUseCase := _cartUsecase.NewCartUsecase(cartRepo, timeoutContext)
	cartController := _cartController.NewCartController(cartUseCase)
	e := echo.New()
	routesInit := _route.ControllerList{
		UserController:     *userController,
		ProductContrller:   *productController,
		CategoryController: *categoryController,
		CartController:     *cartController,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
