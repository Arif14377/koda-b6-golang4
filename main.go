package main

import (
	"fmt"
	lib "koda-b6-golang4/lib"
	"os"
)

func main() {
	defer func() {
		fmt.Println(recover())
		os.Exit(0)
	}()

	pilihan := []string{"Fahrenheit", "Reamur", "Kelvin"}
	var targetPilihan int

	fmt.Println("Masukkan target konversi: (tulis angka)")
	for i := range pilihan {
		fmt.Println(i+1., pilihan[i])
	}
	fmt.Scanf("%d", &targetPilihan)
	target, err := lib.PilihanTarget(targetPilihan)
	if err != nil {
		panic(err)
	}
	suhuAkhir := lib.Conversion(target)
	fmt.Printf("Hasil konversi suhu: %f\n", suhuAkhir)
}
