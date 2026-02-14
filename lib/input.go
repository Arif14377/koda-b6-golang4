package lib

import "fmt"

func pilihanSuhu() {
	pilihan := []string{"Fahrenheit", "Reamur", "Kelvin"}
	var targetPilihan int

	fmt.Println("Masukkan target konversi: (tulis angka)")
	for i := range pilihan {
		fmt.Println(i+1., pilihan[i])
	}
	fmt.Scanf("%d", &targetPilihan)
}

func inputSuhu() {

}
