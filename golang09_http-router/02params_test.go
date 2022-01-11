package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
Params

- httprouter.Handle memiliki parameter yang ketiga, yaitu Params. Untuk apa kegunaan Params?
- Params merupakan tempat untuk meyimpan parameter yang dikirim dari client
- Namun Params ini bukan query parameter, melainkan paramter di URL
- kadang kita butuih membuat URL yang tidak fix, alias bisa berubah-ubah. Misalnya /products/1, /products/2 dan seterusnya
- ServeMux tidak mendukung hal tersebut, namun Router mendukung hal tersebut
- Parameter yang dinamis yang terdapat di URL, secara otomatis dikumpulkan di Params
- Namun, agar Router tahu, kita harus memberi tahu ketika menambahkan Route, dibagian mana kita akan buat URL path nya
  menjadi dinamis
*/

func TestParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
