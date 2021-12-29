package routes

import (
	"acp14/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// Init Echo
	e := echo.New()

	// Routing
	v1 := e.Group("/api/v1/")
	v1.GET("users", controllers.GetUserController)
	v1.POST("product", controllers.CreateProduct)
	v1.GET("products", controllers.GetProductsAll)
	e.GET("products/:id", controllers.GetDetailProduct)

	return e
}
