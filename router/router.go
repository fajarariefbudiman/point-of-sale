package router

import (
	"html/template"
	"io"
	"net/http"

	"pos-echo/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() *echo.Echo {
	/* inisisai object */
	e := echo.New()

	// middleware logger
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/* statis file */
	e.Static("/assets", "assets")

	t := &Template{
		templates: template.Must(template.ParseGlob("./template/*.gohtml")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return t.Render(c.Response().Writer, "index.gohtml", map[string]interface{}{
			"title": "Home Title",
			"body":  "Body Description",
		}, c)
	})
	e.GET("/products", func(c echo.Context) error {
		products := []models.Products{
			{Name: "Product 1", Price: 57.00},
			{Name: "Product 2", Price: 59.00},
		}
		return c.Render(http.StatusOK, "products.gohtml", products)
	})

	return e
}
