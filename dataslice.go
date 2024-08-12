package main

import "fmt"

func main() {

	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret", "April",
		"mei", "juni", "juli",
	}

	var slice = bulan[2:5]

	fmt.Println(bulan)
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(bulan[0])
	fmt.Println(cap(slice))

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	daySlice1[0] = "Saturday"
	daySlice1[1] = "Sunday"

	fmt.Println(days)

	daySlice2 := append(daySlice1, "Weekend") //menambah data baru pada slice dan membuat array baru karna kapasitasnya sudah dtiak muat (append artinya ke replace)
	daySlice2[0] = "yEY!!"                    //mengganti index 0 dari dataSlice2 atau data baru menjadi Yey

	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 4, 5)
	newSlice[0] = "Ana"
	newSlice[1] = "Jannatu"

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	newSlice[3] = "Tset"
	fmt.Println(newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
