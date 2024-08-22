package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, " my name is", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huuuu from", a.Name)
}

func main() {
	var ana Customer
	ana.Name = "Ana"
	ana.Address = "Sagulung"
	ana.Age = 24

	fmt.Println(ana)

	dira := Customer{
		Name:    "Dira",
		Address: "Piayu",
		Age:     23,
	}

	fmt.Println(dira)

	una := Customer{"Una", "Kalimantan", 15}

	fmt.Println(una)

	ana2 := Customer{Name: "Ana"}
	ana2.sayHi("Dira")

	ana2.sayHuu()
	fmt.Println(ana2)

}
