package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = nil
	fmt.Println(person)

	ana := NewMap("")
	if ana == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(ana)
	}
	// fmt.Println(ana)

}
