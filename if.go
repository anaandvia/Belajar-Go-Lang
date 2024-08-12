package main

import "fmt"

func main() {

	name := "Ana"

	if name == "Ana" {
		fmt.Println("Hello " + name)
	} else if name == "Dina" {
		fmt.Println("Hello Girl " + name)
	} else {
		fmt.Println("Hello Bro")
	}

	if length := len(name); length > 6 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")

	}
}
