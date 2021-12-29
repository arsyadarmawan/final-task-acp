package controllers

import (
	"acp14/configs"
	"acp14/models"

	// "acp14/models"

	// "acp14/models"x
	"acp14/models/products"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateProduct(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	// price, err := strconv.Atoi(price)
	var product products.Product
	result := configs.DB.Create(&products.Product{
		Name:       name,
		Desription: description,
		Price:      200,
	})

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Success: false,
			Message: "Error when saving",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Success: true,
		Message: "Success saved data in database",
		Data:    product,
	})
}

func GetProductsAll(c echo.Context) error {
	var products []products.Product
	result := configs.DB.Find(&products)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNoContent, models.BaseResponse{
				Success: true,
				Message: "Empty get data",
				Data:    products,
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
		Data:    products,
	})
}

func GetDetailProduct(c echo.Context) error {
	id := c.Param("id")
	var product products.Product
	result := configs.DB.First(&product, id)

	// return json
	return c.JSON(http.StatusOK, models.BaseResponse{
		Success: true,
		Message: "Success get data in database",
		Data:    result,
	})
}

// func updateProduct(c echo.Context) error {
// 	id := c.Param("id")
// 	name := c.FormValue("name")
// 	description := c.FormValue("description")
// 	var product products.Product
// 	result := configs.DB.First(&product, id).Update({

// 	})

// }
