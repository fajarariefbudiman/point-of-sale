package testing_test

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ProductHTML(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../template/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "JANGAN LUPA MAKAN")
}

func TestProducts(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1323", nil)
	recorder := httptest.NewRecorder()
	ProductHTML(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
