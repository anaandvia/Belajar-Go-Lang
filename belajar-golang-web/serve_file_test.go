package belajargolangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServerFile(w http.ResponseWriter, r *http.Request) {
	// ada proses parsing terlebih dulu

	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}

}

func TestServerFile(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: http.HandlerFunc(ServerFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//go:embed  resources/ok.html
var resourcesOk string

//go:embed  resources/notfound.html
var resourcesNotFound string

func ServerFileEmbed(w http.ResponseWriter, r *http.Request) {
	// ada proses parsing terlebih dulu

	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourcesOk)
	} else {
		fmt.Fprint(w, resourcesNotFound)
	}

}

func TestServerFileEmbed(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: http.HandlerFunc(ServerFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
