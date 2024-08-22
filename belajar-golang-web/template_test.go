package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHtml(w http.ResponseWriter, r *http.Request) {
	templateTest := "<html><body{{ . }}</body><html>"
	// t, err := template.New("SIMPLE").Parse(templateTest)

	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateTest))

	t.ExecuteTemplate(w, "SIMPLE", "Hello everyone")

}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func SimpleHtmlFiles(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Hello everyone")

}

func TestSimpleHtmlFiles(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	SimpleHtmlFiles(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func SimpleHtmlDirectory(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Hello everyone")

}

func TestSimpleHtmlDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	SimpleHtmlDirectory(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(w, "simple.gohtml", "Hello everyone")

}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
