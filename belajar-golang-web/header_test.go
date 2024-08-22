package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// menangkap header
func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8082/", nil)

	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)

	responese := recorder.Result()
	body, _ := io.ReadAll(responese.Body)

	fmt.Println(string(body))
}

// w itu writer, r itu request

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Programmer Zaman Now")
	fmt.Fprint(w, "OK")

	// header tidak sensitif huruf kecil atau besar
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8082/", nil)

	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()
	ResponseHeader(recorder, request)

	responese := recorder.Result()
	body, _ := io.ReadAll(responese.Body)

	fmt.Println(string(body))
	fmt.Println(responese.Header.Get("x-powered-by"))

}
