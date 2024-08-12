package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "ana",
		"address": "sagulung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "Programmer"
	person["name"] = "Ana"

	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Test"
	book["ups"] = "Salah"

	delete(person, "address")
	fmt.Println(person)
	fmt.Println(book)

}
