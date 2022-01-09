package carts

import (
	_cartsDomain "acp14/business/carts"
	_controllers "acp14/controllers"
	_request "acp14/controllers/carts/request"
	_response "acp14/controllers/carts/response"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	usecase _cartsDomain.Usecase
}

func NewCartController(productUsecase _cartsDomain.Usecase) *CartController {
	return &CartController{
		usecase: productUsecase,
	}
}

func (controller *CartController) GetCarts(c echo.Context) error {
	ctx := c.Request().Context()
	carts, err := controller.usecase.GetCarts(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(carts))
}

func (controller *CartController) CreateCart(c echo.Context) error {
	var data _request.CartRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return err
	}
	dataDomain := _cartsDomain.Domain{
		Name:      data.Name,
		Total:     data.Total,
		Price:     uint(data.Price),
		ProductId: uint(data.ProductId),
	}

	fmt.Println(dataDomain)
	_, error := controller.usecase.CreateCart(ctx, dataDomain)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, dataDomain)
}

func (controller *CartController) DeleteCart(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, error := controller.usecase.DeleteCart(ctx, id)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, true)
}

func (controller *CartController) UpdateCart(c echo.Context) error {
	var request _request.UpdateRequest
	ctx := c.Request().Context()
	if err := c.Bind(&request); err != nil {
		return err
	}

	dataDomain := _cartsDomain.Domain{
		Id: request.Id,
		// Name:       request.Name,
		// Price:     uint(request.Price),
		// ProductId: request.ProductId,
		Total: request.Total,
	}

	fmt.Println(dataDomain)
	product, err := controller.usecase.UpdateCart(ctx, dataDomain)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	return _controllers.NewSuccessResponse(c, _response.FromDomain(product))
}
