package main

import "fmt"

// defer, panic, recover
func logging() {
	fmt.Println("Selesai memanggil aplikasi")
}

func endApp() {
	// recover selalu di simpan di defer function
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message : ", message)
	}

	fmt.Println("End App")

}

func runApp(error bool) {
	// defer logging() //defer akan dijalankan setelah func runApp walaupun si runApp mengalami error akan ttp di eksekusi
	// fmt.Println("Run Aplikasi")
	defer endApp()
	if error {
		panic("Error App")
	}

	fmt.Println("Aplikasi berjalan")

}

func main() {
	runApp(true)
	fmt.Println("Hallo")

}
