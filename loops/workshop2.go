package loops

import "fmt"

func Workshop2() {
	sayi := 0
	sayac := 0

	fmt.Print("Bir Sayı Giriniz: ")
	fmt.Scanln(&sayi)

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			sayac++
		}
	}
	if sayac != 0 {
		fmt.Printf("%v: Asal değildir", sayi)
	} else {
		fmt.Printf("%v: Asaldır", sayi)
	}
}
