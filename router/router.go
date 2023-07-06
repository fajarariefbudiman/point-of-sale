package router

import (
	"pos-echo/controllers"
	"pos-echo/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	/* inisisai object */
	e := echo.New()
	_ = db.NewDB()
	// middleware logger
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", controllers.HomeController)
	e.GET("/products", controllers.ProductsControllers)
	e.GET("/products/:slug", controllers.FindProductsBySlugControllers)
	e.GET("/carts", controllers.GetCart)
	e.POST("/carts", controllers.AddItemCart)
	return e
}
