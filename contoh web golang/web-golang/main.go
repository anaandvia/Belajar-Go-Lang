package main

import (
	"log"
	"net/http"
	"web-golang/config"
	"web-golang/controllers/categoriescontroller"
	"web-golang/controllers/homecontroller"
)

func main() {
	config.ConnectDB()

	// 1. HomeController
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categoriescontroller.Index)
	http.HandleFunc("/categories/add", categoriescontroller.Add)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	log.Println("server running")
	http.ListenAndServe(":8082", nil)

}
