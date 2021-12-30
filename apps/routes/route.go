package routes

// import (
// 	"acp14/controllers"

// 	"github.com/labstack/echo/v4"
// )

// func New() *echo.Echo {
// 	// Init Echo
// 	e := echo.New()

// 	// Routing
// 	v1 := e.Group("/api/v1/")
// 	v1.GET("users", controllers.GetUserController)
// 	v1.POST("product", controllers.CreateProduct)
// 	v1.GET("products", controllers.GetProductsAll)
// 	e.GET("products/:id", controllers.GetDetailProduct)

// 	return e
// }

import (
	_products "acp14/controllers/products"
	_users "acp14/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController   _users.UserController
	ProductContrller _products.ProductController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	c.GET("users", cl.UserController.GetUser)
	c.POST("users", cl.UserController.Register)
	c.GET("products", cl.ProductContrller.GetProduct)
	c.POST("product", cl.ProductContrller.CreateProduct)
	c.DELETE("products/:id", cl.ProductContrller.DeleteProduct)

}
