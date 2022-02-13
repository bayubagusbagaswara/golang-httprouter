package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {

	// buat routernya
	router := httprouter.New()

	// buat httpmethod dan handle untuk routenya
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "Hello World")
	})

	// bikin requestnya, masukkan http method nya
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	// bikin response nya
	recorder := httptest.NewRecorder()

	// jalankan routernya menggunakan ServeHttp
	router.ServeHTTP(recorder, request)

	// tangkap responsenya
	response := recorder.Result()
	b, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(b))
}
