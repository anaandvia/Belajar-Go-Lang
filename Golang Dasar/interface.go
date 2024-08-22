package main

import "fmt"

// fmt.Println(a ...interface[])
type HasName interface {
	getName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var ana2 Person
	ana2.Name = "Ana"
	ana := Person{Name: "Ana"}
	sayHello(ana2)
	sayHello(ana)
	pus := Person{Name: "Meow"}
	sayHello(pus)

	ups := Ups(5)
	fmt.Println(ups)

}
