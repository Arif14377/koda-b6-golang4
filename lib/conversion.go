package lib

import (
	"errors"
	"fmt"
)

func PilihanTarget(target int) (int, error) {
	if target > 3 || target < 1 {
		err := errors.New("Pilihan target konversi tidak tersedia.")
		return target, err
	}
	return target, nil
}

func Conversion(target int) float64 {
	var suhuCelcius float64
	var suhuAkhir float64

	fmt.Println("# KONVERSI SUHU CELCIUS")
	fmt.Printf("%s", "Masukkan suhu: ")
	fmt.Scanf("%f", &suhuCelcius)

	switch target {
	case 1:
		toFahrenheit := (suhuCelcius * 9.0 / 5.0) + 32.0
		suhuAkhir = toFahrenheit
	case 2:
		toReamur := suhuCelcius * 4.0 / 5.0
		suhuAkhir = toReamur
	case 3:
		toKelvin := suhuCelcius + 273.0
		suhuAkhir = toKelvin
	}

	return suhuAkhir
}
