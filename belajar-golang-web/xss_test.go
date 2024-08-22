package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// text/template itu tidak bisa melakukan auto escape jadi gunakan html/template

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  "<p> test <script> alert('anda di hack') </script></p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TemplateAutoEscapeDisable(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  template.HTML("<p> test</p>"),
	})
}

func TestTemplateAutoEscapeDisable(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisable(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeDisableServer(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisable),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Go-Lang Auto Escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082?body=<p>alert</p>", nil)

	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
