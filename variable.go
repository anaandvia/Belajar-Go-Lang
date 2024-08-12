package main

import "fmt"

func main() {
	var name string
	var test bool

	name = "Ana Jannatu Uzlifat"
	fmt.Println(name)

	name = "Ana Jannatu"
	fmt.Println(name)

	test = true
	fmt.Println(test)

	var name2 = "Ana Jannatu Uzlifat"
	fmt.Println(name2)

	var age int8 = 24
	fmt.Println(age)

	name1 := "Ana Jannatu Uzlifat"

	fmt.Println(name1)

	name1 = "Popo"
	fmt.Println(name1)

	var (
		firstname = "test"
		lastname  = "test2"
	)

	fmt.Println(firstname + " " + lastname)

}
