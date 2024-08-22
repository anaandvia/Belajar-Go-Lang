package main

import "fmt"

func sumAll(numbers ...int) int {
	// ... menandakan bisa memaasukan banyak variabel dengan tipe data int
	// hanya 1 varagrs 1 saja, dan hanya bisa di akhir paramwter

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	//varargs = variabel argument
	total := sumAll(10, 10, 30)
	fmt.Println(total)

	numbers := []int{10, 20, 30}
	total = sumAll(numbers...)
	fmt.Println(total)

}
