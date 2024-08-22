package main

import "fmt"

func sayhello() {
	fmt.Println("Hello")
}

func getHello(name string) string {
	if name == "" {
		return "Hello guest"
	} else {
		return "Hello " + name
	}

}

func main() {
	sayhello()

	name2 := "Ana"
	result := getHello(name2)
	fmt.Println(result)
}
