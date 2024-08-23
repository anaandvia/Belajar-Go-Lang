package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// panic("Ups")
		fmt.Fprint(w, "Hello Http Router")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8082",
	}

	server.ListenAndServe()
}
