package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddress(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	address := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := &address //pointer &

	address2.City = "Subang"
	address4 := new(Address)

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}
	// siapaun yng mengacu ke addres 1 maka akan berubah jika menggunakan tanda bintang
	fmt.Println(address)
	fmt.Println(address2)
	fmt.Println(address4)

	alamat := Address{
		City:     "Subang",
		Province: "Jabar",
		Country:  "",
	}

	alamat2 := &alamat
	changeAddress(&alamat)
	// changeAddress(alamat2)

	fmt.Println(alamat)
	fmt.Println(alamat2)

}
