package main

import "fmt"

func main() {

	type NoKTP string

	var Ktp NoKTP = "12345678"

	fmt.Println(Ktp)
}
