package categories

import (
	_categoriesDomain "acp14/business/categories"
	_controllers "acp14/controllers"
	_request "acp14/controllers/categories/request"
	_response "acp14/controllers/categories/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	usecase _categoriesDomain.Usecase
}

func NewCategoryController(productUsecase _categoriesDomain.Usecase) *CategoryController {
	return &CategoryController{
		usecase: productUsecase,
	}
}

func (controller *CategoryController) GetCategories(c echo.Context) error {
	ctx := c.Request().Context()
	products, err := controller.usecase.GetCategories(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(products))
}

func (controller *CategoryController) CreateCategory(c echo.Context) error {
	var data _request.ProductRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return err
	}
	dataDomain := _categoriesDomain.Domain{
		Name: data.Name,
	}
	_, error := controller.usecase.CreateCategory(ctx, dataDomain)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, dataDomain)
}

func (controller *CategoryController) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, error := controller.usecase.DeleteCategory(ctx, id)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, true)
}
