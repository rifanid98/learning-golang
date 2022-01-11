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
Router

- Inti dari library HttpRouter adalah struct Router
- Router ini merupakan implementasi dari http.Handler, sehingga kita bisa dengan mudah menambahkan ke dalam http.Server
- Untuk membuat Router, kita bisa menggunakan function httprouter.New(), yang akan mengembalikan Router pointer

HTTP Method

- Router mirip dengan ServeMux, dimana kita bisa menambahkan route ke dalam Router
- Kelebihan dibandingakn dengan SeverMux adalah, pada Router kita bisa menentukan HTTP Method yang ingin kita gunakan.
  Misalnya seperti GET, POST, PUT dan lain-lain
- Cara menambahkan route ke dalam Router adalah dengan menggunakan function yang sama dengan HTTP Methodnya. Misal
  router.GET(), router.POSTT() dan lain-lain

httprouter.Handler

- Saat kita menggunakan ServeMux, ketika menambahkan route, kita bisa menambahkan http.Handler
- Berbeda dengan Router, pada Router kita tidak menggunakan http.Handler lagi, melainkan menggunakan type httprouter.Handle
- Perbedaan dengan http.Handler adalah, pada httprouter.Handle, terdapat parameter ketiga yaitu Params

*/

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")

	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))
}
