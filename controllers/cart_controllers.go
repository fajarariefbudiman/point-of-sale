package controllers

import (
	"net/http"
	"pos-echo/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCart(c echo.Context) error {

}

func AddItemCart(c echo.Context) error {
	productId := c.FormValue("product_Id")
	strId, _ := strconv.Atoi(productId)
	quantity, _ := strconv.Atoi(c.FormValue("quantity"))
	product, err := models.FindProductsById(strId)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "product"+product.Slug)
	}
	if quantity > product.Stock {
		return c.Redirect(http.StatusSeeOther, "product"+product.Slug)
	}

}
