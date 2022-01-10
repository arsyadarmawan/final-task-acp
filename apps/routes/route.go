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

	j := c.Group("api/")

	j.GET("users", cl.UserController.GetUser)
	j.POST("users", cl.UserController.Register)
	j.POST("login", cl.UserController.Login)

	j.GET("products", cl.ProductContrller.GetProduct)
	j.GET("product/:id", cl.ProductContrller.GetProductById)
	j.GET("category/products/:id", cl.ProductContrller.SearchCategoy)
	j.PUT("product/:id", cl.ProductContrller.UpdateProduct)
	j.POST("product", cl.ProductContrller.CreateProduct)
	j.DELETE("product/:id", cl.ProductContrller.DeleteProduct)

	j.GET("categories", cl.CategoryController.GetCategories)
	j.POST("category", cl.CategoryController.CreateCategory)
	j.DELETE("category/:id", cl.CategoryController.DeleteCategory)

	j.GET("carts", cl.CartController.GetCarts)
	j.POST("cart", cl.CartController.CreateCart)
	j.DELETE("cart/:id", cl.CartController.DeleteCart)

	// j := c.Group("/access_token")
	// j.Use(middleware.JWT([]byte("ThisisSecretGais")))
}
