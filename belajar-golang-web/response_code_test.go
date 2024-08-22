package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name is empty")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponseCoseInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://loaclhost:8082", nil)

	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCoseValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://loaclhost:8082/?name=ana", nil)

	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
