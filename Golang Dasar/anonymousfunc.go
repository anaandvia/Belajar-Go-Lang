package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked")
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {

	blacklist := func(name string) bool {
		return name == "Admin"
	}

	registerUser("Admin", blacklist)
	registerUser("Ana", blacklist)

}
