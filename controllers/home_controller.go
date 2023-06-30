package controllers

import (
	"io"
	"pos-echo/models"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Templateee struct {
	templates *template.Template
}

func (t *Templateee) Renderrr(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func HomeController(c echo.Context) error {
	t := &Template{
		templates: template.Must(template.ParseGlob("./template/*.gohtml")),
	}
	c.Echo().Renderer = t
	c.Echo().Static("/assets", "assets") // Serve static files from the "assets" directory
	perPage := 8
	page := 1
	products, _, err := models.GetProducts(perPage, page)
	if err != nil {
		// Tangani error jika terjadi
		return err
	}

	return t.Render(c.Response().Writer, "index.gohtml", map[string]interface{}{
		"products": products,
	}, c)
}
