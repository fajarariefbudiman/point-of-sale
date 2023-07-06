package testing

import (
	"net/http"
	"net/http/httptest"
	"pos-echo/controllers"
	"pos-echo/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetShopingCartId(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Panggil fungsi GetShopingCartId
	shopingCartID := controllers.GetShopingCartId(c)

	// Periksa apakah ID keranjang belanja telah diset dengan benar
	assert.NotEmpty(t, shopingCartID)
}

func TestGetCart(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Panggil fungsi GetCart
	err := controllers.GetCart(c)

	// Periksa apakah tidak ada kesalahan yang terjadi
	assert.NoError(t, err)

	// Periksa apakah output yang diharapkan sesuai dengan yang sebenarnya
	expectedOutput := "cart-id ==> <ID keranjang belanja>"
	assert.Equal(t, expectedOutput, rec.Body.String())
}

type MockCartModel struct{}

func (m *MockCartModel) GetCart(cartId string) (*models.Cart, error) {
	// Simulasikan implementasi GetCart
	// Anda dapat mengganti dengan logika pengujian yang sesuai
	// Misalnya, mengembalikan keranjang belanja yang ada atau nil jika tidak ada
	exisCart := &models.Cart{
		// ...
	}
	return exisCart, nil
}

func (m *MockCartModel) CreateCart(cartId string) (*models.Cart, error) {
	// Simulasikan implementasi CreateCart
	// Anda dapat mengganti dengan logika pengujian yang sesuai
	// Misalnya, mengembalikan keranjang belanja yang baru dibuat
	newCart := &models.Cart{
		// ...
	}
	return newCart, nil
}

func TestGetShopingCart(t *testing.T) {
	// Persiapkan objek CartModel palsu

	// Panggil fungsi GetShopingCart dengan objek CartModel palsu

	exisCart, err := controllers.GetShopingCart("cart-123")

	// Periksa apakah tidak ada kesalahan yang terjadi
	assert.NoError(t, err)

	// Periksa apakah objek Cart yang dikembalikan sesuai dengan yang diharapkan
	expectedCart := &models.Cart{
		// ...
	}
	assert.Equal(t, expectedCart, exisCart)
}
