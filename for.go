package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)

	}

	slice := []string{"Ana", "Jannatu", "Uzlifat"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for i, name := range slice {
	// 	fmt.Println("index", i, "=", name)
	// }

	for _, name := range slice {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "ana"
	person["title"] = "narnia"

	for key, name := range person {
		fmt.Println(key, "=", name)
	}
}
