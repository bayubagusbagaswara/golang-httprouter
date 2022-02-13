package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// bikin router dengan New
	// httprouter adalah turunan/implementasi dari handler itu sendiri
	router := httprouter.New()

	// http method lalu masukkan URL nya dan handle function nya
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "Hello HttpRouter")
	})

	// bikin server
	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	// jalankan server nya
	server.ListenAndServe()
}
