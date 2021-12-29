package controllers

import (
	"acp14/helpers"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		StatusCode int      `json:"statusCode"`
		Message    string   `json:"message"`
		Messages   []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.StatusCode = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, err error) error {
	response := BaseResponse{}

	status := helpers.GenerateErrorCode(err)

	response.Meta.StatusCode = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}
