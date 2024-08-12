package main

import "fmt"

func main() {

	var a = 10
	var b = 3
	var c = 14

	var d = a + b - c
	var e = a*b - c

	fmt.Println(d)
	fmt.Println(e)

	a += 10 // ini sama dengan a = a+10
	fmt.Println(a)

	a--
	fmt.Println(a)

	var a2 = -200
	fmt.Println(a2)

}
