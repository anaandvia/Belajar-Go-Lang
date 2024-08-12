package main

import "fmt"

func main() {
	var (
		nilaiAkhir = 80
		absensi    = 60

		lulusNilaiAKhir = nilaiAkhir > 70
		lulusAbsensi    = absensi > 70

		lulus = lulusNilaiAKhir || lulusAbsensi
	)

	fmt.Println(lulus)

}
