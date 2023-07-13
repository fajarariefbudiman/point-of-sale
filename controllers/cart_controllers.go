package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"pos-echo/models"
	"strconv"
	"text/template"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type TemplateCart struct {
	templates *template.Template
}

func (t *TemplateCart) RenderCart(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var sessionShoppingCart = "shopping-cart-session"

func GetShopingCartId(c echo.Context) string {
	session, _ := store.Get(c.Request(), sessionShoppingCart)
	if session.Values["cart-id"] == nil {
		session.Values["cart-id"] = uuid.New().String()
		_ = session.Save(c.Request(), c.Response())
	}
	return fmt.Sprintf("%s", session.Values["cart-id"])
}

func GetShopingCart(cartId string) (*models.Cart, error) {

	exisCart, err := models.GetCart(cartId)
	if err != nil {
		exisCart, err = models.CreateCart()
		if err != nil {
			return nil, err
		}
	}

	exisCart.CalcuCart(cartId)

	return exisCart, nil

}

func GetCart(c echo.Context) error {
	//render
	t := &Template{
		templates: template.Must(template.ParseGlob("./template/*.gohtml")),
	}
	c.Echo().Renderer = t
	c.Echo().Static("/assets", "assets")

	var cart *models.Cart
	cartId := GetShopingCartId(c)
	cart, _ = GetShopingCart(cartId)

	return t.Render(c.Response().Writer, "cart.gohtml", map[string]interface{}{
		"cart": cart,
	}, c)

}

func AddItemCart(c echo.Context) error {
	productId := c.FormValue("product_id")
	strId, _ := strconv.Atoi(productId)
	quantity, _ := strconv.Atoi(c.FormValue("quantity"))
	product, err := models.FindProductsById(strId)
	if err != nil {
		// Tangani kesalahan saat mencari produk
		fmt.Println("Error finding product:", err)
		return c.Redirect(http.StatusSeeOther, "/product/"+product.Slug)
	}
	if quantity > product.Stock {
		// Tangani kesalahan jika kuantitas melebihi stok produk
		return c.Redirect(http.StatusSeeOther, "/product/"+product.Slug)
	}

	var cart *models.Cart
	cartId := GetShopingCartId(c)
	cart, err = GetShopingCart(cartId)
	if err != nil {
		// Tangani kesalahan saat mendapatkan keranjang belanja
		fmt.Println("Error getting shopping cart:", err)
		// Anda dapat memberikan respons yang sesuai kepada pengguna
		return err
	}

	_, err = cart.AddItem(models.Cart_Item{
		Product_Id: productId,
		Quantity:   quantity,
	})
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/product/"+product.Slug)
	}

	return c.Redirect(http.StatusSeeOther, "/carts")
}

func UpdateCart(c echo.Context) error {
	cartId := GetShopingCartId(c)
	cart, err := GetShopingCart(cartId)
	if err != nil {
		fmt.Println("Error getting shopping cart:", err)
		return err
	}
	for _, item := range cart.Cart_Items {
		quantity, _ := strconv.Atoi(c.FormValue(item.Id))
		_, err := models.UpdateQuantity(item.Id, quantity)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/carts")
		}
	}

	return c.Redirect(http.StatusSeeOther, "/carts")

}
