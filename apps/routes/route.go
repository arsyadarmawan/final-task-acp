package routes

import (
	_carts "acp14/controllers/carts"
	_categories "acp14/controllers/categories"
	_products "acp14/controllers/products"

	_users "acp14/controllers/users"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController     _users.UserController
	ProductContrller   _products.ProductController
	CategoryController _categories.CategoryController
	CartController     _carts.CartController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {

	c.GET("users", cl.UserController.GetUser)
	c.POST("users", cl.UserController.Register)

	c.GET("products", cl.ProductContrller.GetProduct)
	c.GET("product/:id", cl.ProductContrller.GetProductById)
	c.PUT("product/:id", cl.ProductContrller.UpdateProduct)
	c.POST("product", cl.ProductContrller.CreateProduct)
	c.DELETE("product/:id", cl.ProductContrller.DeleteProduct)

	c.GET("categories", cl.CategoryController.GetCategories)
	c.POST("category", cl.CategoryController.CreateCategory)
	c.DELETE("category/:id", cl.CategoryController.DeleteCategory)

	c.GET("carts", cl.CartController.GetCarts)
	c.POST("cart", cl.CartController.CreateCart)
	c.DELETE("cart/:id", cl.CartController.DeleteCart)

	// j := c.Group("/access_token")
	// j.Use(middleware.JWT([]byte("ThisisSecretGais")))
}
