package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
Serve File

- Pada materi Go-Lang Web, kita sudah pernah membahas tentang Serve File
- Pada Router pun, mendukung serve static file menggunakan function ServeFiles(Path, FileSystem)
- Dimana pada Path, kita harus menggunakan Catch All Parameter
- Sedangkan pada FileSystem kita bisa melakukan manual load dari folder atau menggunakan golang embed, seperti yang
  pernah kita bahas di materi Go-Lang Web
*/

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))
}
