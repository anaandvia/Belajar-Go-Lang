package main

import (
	"fmt"
)

// typedeclaration

type Filter func(string) string

func sayHelloFilter(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("Hello ", nameFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	filter := spamFilter
	sayHelloFilter("Anjing", filter)
	sayHelloFilter("Ana", filter)

}
