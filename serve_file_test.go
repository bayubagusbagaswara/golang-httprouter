package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {

	router := httprouter.New()

	// bikin sub sistem dulu, masuk ke directory nya dulu
	directory, _ := fs.Sub(resources, "resources")

	// nama untuk catch all parameternya harus *filepath, tidak boleh yang lain
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)

	recorder := httptest.NewRecorder()

	// jalankan router nya
	router.ServeHTTP(recorder, request)
	r := recorder.Result()
	b, _ := io.ReadAll(r.Body)

	assert.Equal(t, "Hello HttpRouter", string(b))
}
