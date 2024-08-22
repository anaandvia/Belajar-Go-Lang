package categoriescontroller

import (
	"net/http"
	"text/template"
	"web-golang/models/categoriesmodel"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categoriesmodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/categories/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
