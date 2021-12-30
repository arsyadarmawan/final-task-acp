package products

import (
	_productsDomain "acp14/business/products"
	_controllers "acp14/controllers"
	_request "acp14/controllers/products/request"
	_response "acp14/controllers/products/response"
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
	return _controllers.NewSuccessResponse(c, "sss")
}

func (controller *ProductController) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, error := controller.usecase.DeleteProduct(ctx, id)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, "sss")
}
