package main

import (
	_userUsecase "acp14/business/users"
	_userController "acp14/controllers/users"
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
	e := echo.New()
	routesInit := _route.ControllerList{
		UserController: *userController,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
