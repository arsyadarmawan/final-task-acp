package controllers

import (
	"acp14/configs"
	"acp14/models"
	"acp14/models/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUserController(c echo.Context) error {
	// get user from db
	var users []users.User
	result := configs.DB.Find(&users)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNoContent, models.BaseResponse{
				Success: true,
				Message: "Empty get data",
				Data:    users,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, models.BaseResponse{
				Success: false,
				Message: "Failed get data in database",
				Data:    nil,
			})
		}
	}

	// return json
	return c.JSON(http.StatusOK, models.BaseResponse{
		Success: true,
		Message: "Success get data in database",
		Data:    users,
	})
}
