package products

import (
	_productsDomain "acp14/business/products"
	_controllers "acp14/controllers"
	_request "acp14/controllers/products/request"
	_response "acp14/controllers/products/response"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	usecase _productsDomain.Usecase
}

func NewProductController(productUsecase _productsDomain.Usecase) *ProductController {
	return &ProductController{
		usecase: productUsecase,
	}
}

func (controller *ProductController) GetProduct(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := controller.usecase.GetProducts(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(products))
}

func (controller *ProductController) SearchCategoy(c echo.Context) error {
	ctx := c.Request().Context()
	productRequest := _request.CategoryId{}
	c.Bind(&productRequest) // error handling

	value := c.Param("id")
	number, err := strconv.ParseUint(value, 10, 32)

	IdCategory := _request.CategoryId{
		Id: uint(number),
	}

	products, err := controller.usecase.SearchCategoy(ctx, int(IdCategory.Id))

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(products))
}

func (controller *ProductController) CreateProduct(c echo.Context) error {
	var data _request.ProductRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return err
	}
	dataDomain := _productsDomain.Domain{
		Name:       data.Name,
		Price:      uint(data.Price),
		CategoryId: uint(data.CategoryId),
	}
	_, error := controller.usecase.CreateProduct(ctx, dataDomain)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, dataDomain)
}

func (controller *ProductController) DeleteProduct(c echo.Context) error {
	var request _request.DeleteProduct
	ctx := c.Request().Context()
	if err := c.Bind(&request); err != nil {
		return err
	}
	Product, err := controller.usecase.DeleteProduct(ctx, int(request.Id))

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Product))
}

func (controller *ProductController) GetProductById(c echo.Context) error {
	ctx := c.Request().Context()
	productRequest := _request.ProductId{}
	c.Bind(&productRequest) // error handling

	value := c.Param("id")
	number, err := strconv.ParseUint(value, 10, 32)

	finalNumber := _request.ProductId{
		Id: uint(number),
	}

	Product, err := controller.usecase.GetProductById(ctx, int(finalNumber.Id))

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.FromDomain(Product))
}

func (controller *ProductController) UpdateProduct(c echo.Context) error {
	var request _request.UpdateProduct
	ctx := c.Request().Context()
	if err := c.Bind(&request); err != nil {
		return err
	}

	dataDomain := _productsDomain.Domain{
		Id:         request.Id,
		Name:       request.Name,
		Price:      uint(request.Price),
		CategoryId: request.CategoryId,
	}

	fmt.Println(dataDomain)
	product, err := controller.usecase.UpdateProduct(ctx, dataDomain)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	return _controllers.NewSuccessResponse(c, _response.FromDomain(product))
}
