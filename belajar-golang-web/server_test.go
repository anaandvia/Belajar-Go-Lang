package belajargolangweb

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8082",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
