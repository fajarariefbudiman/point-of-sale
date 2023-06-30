package controllers

import (
	"io"
	"pos-echo/models"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Templatee struct {
	templates *template.Template
}

func (t *Templatee) Renderr(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func ProductsControllers(c echo.Context) error {
	t := &Template{
		templates: template.Must(template.ParseGlob("./template/*.gohtml")),
	}
	c.Echo().Renderer = t
	c.Echo().Static("/assets", "assets")
	q := c.Request().URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}
	perPage := 9
	products, totalrows, err := models.GetProducts(perPage, page)
	if err != nil {
		return err
	}
	pagination, _ := GetPagintionLink(PaginationParam{
		Path:        "products",
		TotalRows:   totalrows,
		PerPage:     perPage,
		CurrentPage: page,
	})

	return t.Render(c.Response().Writer, "products.gohtml", map[string]interface{}{
		"products":   products,
		"pagination": pagination,
	}, c)
}
