package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HalloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://loaclahost:8082/helo", nil)
	recorder := httptest.NewRecorder()

	HalloHandler(recorder, request)

	respone := recorder.Result()
	body, _ := io.ReadAll(respone.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
