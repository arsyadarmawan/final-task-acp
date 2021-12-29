package users

import (
	_userDomain "acp14/business/users"
	_controllers "acp14/controllers"
	_response "acp14/controllers/users/response"

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
	users, err := controller.usecase.GetUser(ctx)

	if err != nil {
		return _controllers.NewErrorResponse(c, err)
	}

	return _controllers.NewSuccessResponse(c, _response.ToListFromDomain(users))
}
