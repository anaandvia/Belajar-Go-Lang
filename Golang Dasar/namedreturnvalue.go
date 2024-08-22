package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Ana"
	lastName = "Jannatu"

	return
}

func main() {
	//nama variable tidak apa apa jika tidak sama dengan variable di function yang di panggil
	firstName, lastName := getCompleteName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
