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

func TestRouterPatternNamedParameter(t *testing.T) {

	router := httprouter.New()

	// ada 2 parameter yakni id dan itemId
	router.GET("/products/:id/items/:itemId", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id") + " Item " + p.ByName("itemId")
		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/2", nil)
	recorder := httptest.NewRecorder()

	// jalankan routernya
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	b, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 2", string(b))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()

	// tangkap semua parameter setelah URL /images/
	router.GET("/images/*image", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Image : " + p.ByName("image")
		fmt.Fprint(rw, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	// jalankan routernya
	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	b, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(b))
}
