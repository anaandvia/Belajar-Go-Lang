package main

import "fmt"

func main() {

	name := "Ana"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Ana":
		fmt.Println("Hallo Ana")
	default:
		fmt.Println("Hallo Semua")
	}

	switch length := len(name); length > 24 {
	case true:
		fmt.Println("Nama terlau panjang")
	case false:
		fmt.Println("Nama Sudah Benar")

	}

}
