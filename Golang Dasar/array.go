package main

import "fmt"

func main() {
	var name [2]string

	name[0] = "Ana"
	name[1] = "jannatu"

	fmt.Println(name)
	fmt.Println(name[1])

	var values = [4]int{
		100,
		200,
		300,
		400,
	}

	fmt.Println(values)
	fmt.Println(values[1])

	fmt.Println(len(values)) //cek jumlah panjang array bukan datanya, datanya bisa jadi kosong

}
