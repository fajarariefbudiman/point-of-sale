package router

import (
	"html/template"
	"io"
	"pos-echo/controllers"
	"pos-echo/db"
	"pos-echo/models"
	"strconv"

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
	_ = db.NewDB()
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
		pagination, _ := controllers.GetPagintionLink(controllers.PaginationParam{
			Path:        "products",
			TotalRows:   totalrows,
			PerPage:     perPage,
			CurrentPage: page,
		})

		return t.Render(c.Response().Writer, "products.gohtml", map[string]interface{}{
			"products":   products,
			"pagination": pagination,
		}, c)
	})
	return e
}
