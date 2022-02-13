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

func TestNotFound(t *testing.T) {

	router := httprouter.New()

	// apapun yang di request hasilnya not found, karena kita tidak memasukkan method URL di router nya
	// isinya adalah http Handler
	router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Gak ketemu")
	})

	// jika kita kirim request, maka akan not found
	request := httptest.NewRequest("GET", "http://localhost:3000/404", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	r := recorder.Result()
	b, _ := io.ReadAll(r.Body)

	assert.Equal(t, "Gak ketemu", string(b))

}
