package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	// ada proses parsing terlebih dulu
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// bisa gunakan langsung tanpa parsing
	r.PostFormValue("firstname")

	firstname := r.PostForm.Get("firstname")
	lastname := r.PostForm.Get("lastname")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)

}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=ana&lastname=janatu")
	request := httptest.NewRequest(http.MethodPost, "http://loaclhost:8082", requestBody)

	// harus dicanntumkan ketika membuat form post
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
