package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/layout.gohtml", "./templates/footer.gohtml"))

	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Ana",
	})

}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
