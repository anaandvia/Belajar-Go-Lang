package main

import "fmt"

func main() {
	name := "ana"
	counter := 0

	// hati hati pengunaan closure, data di atas bisa di akses ke scop func

	increment := func() {
		name := "Budi" //harus di deklarasikan ulang agar tidak mengakses variable di atas
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)

}
