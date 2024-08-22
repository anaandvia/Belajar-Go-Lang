package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "hello " + name + ", My name is " + myPage.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Budi" }}`))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "ana",
	})

}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse("{{ upper .Name }}"))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "ana",
	})

}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateFunctionPipelines(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	// jika ingin hasil function diberi function lagi hanya perlu tambahkan " | nama functionnya "
	t = template.Must(t.Parse("{{ sayHello .Name | upper }}"))

	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "ana",
	})

}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8082", nil)

	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
