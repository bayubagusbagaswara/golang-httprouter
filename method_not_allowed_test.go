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

func TestMethodNotAllowed(t *testing.T) {

	router := httprouter.New()

	// tambahkan method MethodNotAllowed ke router
	router.MethodNotAllowed = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Gak Boleh")
	})

	// tambahkan route URL method POST ke router
	router.POST("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "POST")
	})

	// bikin request dengan URL yang benar, tapi method HTTP nya berbeda
	// maka akan dihandle oleh MethodNotAllowed
	request := httptest.NewRequest("PUT", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	b, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Boleh", string(b))
}
