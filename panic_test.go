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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	// tambahkan panic handler ke routernya
	// biasanya tujuannya untuk menghandle jika terjadi panic pada router, misal ingin mengubah pesan panic nya
	router.PanicHandler = func(rw http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(rw, "Panic : ", error)
	}

	// tambahkan route ke routernya, harapannya jika route GET "/" ini akan menghasilkan panic
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	// jalankan routernya
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	b, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(b))
}
