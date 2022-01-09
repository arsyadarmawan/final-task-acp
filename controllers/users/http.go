package users

import (
	_userDomain "acp14/business/users"
	_controllers "acp14/controllers"
	_request "acp14/controllers/users/request"
	_response "acp14/controllers/users/response"
	"fmt"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase _userDomain.Usecase
}

func NewUserController(userUsercase _userDomain.Usecase) *UserController {
	return &UserController{
		usecase: userUsercase,
	}
}

func (controller *UserController) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	fmt.Println(ctx)
	users, err := controller.usecase.GetUser(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(users))
}

func (controller *UserController) Register(c echo.Context) error {
	var data _request.UserRequest
	ctx := c.Request().Context()
	if err := c.Bind(&data); err != nil {
		return err
	}

	fmt.Println(data, ctx)
	dataDomain := _userDomain.Domain{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	_, error := controller.usecase.Register(ctx, dataDomain)

	if error != nil {
		return _controllers.NewErrorResponse(c, error)
	}
	return _controllers.NewSuccessResponse(c, dataDomain)
}

func (controller *UserController) Login(c echo.Context) error {
	var data _request.UserRequest
	ctx := c.Request().Context()

	if err := c.Bind(&data); err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	// fmt.Println(data.Email, data.Password)
	token, err := controller.usecase.Login(ctx, data.Email, data.Password)
	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}
	return _controllers.NewSuccessResponse(c, token)
}
