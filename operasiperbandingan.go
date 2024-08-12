package main

import "fmt"

func main() {

	var a = 8
	var c = 6

	var d = a < c

	fmt.Println(d)

	var (
		name1 = "ana"
		name2 = "jannatu"
		name3 = name1 != name2 //hasilnya akan boolean
	)

	fmt.Println(name3)
}
