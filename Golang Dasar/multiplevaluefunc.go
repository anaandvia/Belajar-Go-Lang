package main

import "fmt"

func getFullName() (string, string) {
	return "Ana", "Jannatu"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	firstName1, _ := getFullName()
	fmt.Println(firstName1)
}
