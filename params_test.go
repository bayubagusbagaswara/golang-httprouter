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

func TestParams(t *testing.T) {
	router := httprouter.New()

	// id disini bernilai dinamis dan disimpan dalam Params
	router.GET("/products/:id", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id")
		fmt.Fprint(rw, text)
	})

	// bikin request
	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)

	// bikin recorder untuk response
	recorder := httptest.NewRecorder()

	// jalankan routernya
	router.ServeHTTP(recorder, request)

	// tangkap responsenya
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}
