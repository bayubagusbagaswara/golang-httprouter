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

// bikin struct LogMiddleware
type LogMiddleware struct {
	http.Handler
}

// bikin function yang mengikuti kontrak LogMiddleware
func (middleware *LogMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive request")
	// lalu kita forward ke handle selanjutnya
	middleware.Handler.ServeHTTP(rw, r)
}

func TestMiddleware(t *testing.T) {

	router := httprouter.New()

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "Middleware")
	})

	// bikin middleware, dan masukkan routernya ke handler
	middleware := LogMiddleware{Handler: router}

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	// jalankan middleware
	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	b, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(b))
}
