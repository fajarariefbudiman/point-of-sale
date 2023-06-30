package controllers

import (
	"io"
	"net/http"
	"pos-echo/models"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func FindProductsBySlugControllers(c echo.Context) error {
	slug := c.Param("slug")
	if slug == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	product, err := models.FindProductsBySlug(slug)
	if err != nil {
		return err
	}
	t := &Template{
		templates: template.Must(template.ParseGlob("./template/*.gohtml")),
	}
	c.Echo().Renderer = t
	c.Echo().Static("/assets", "assets")
	return t.Render(c.Response().Writer, "product.gohtml", map[string]interface{}{
		"product": product,
	}, c)
}
