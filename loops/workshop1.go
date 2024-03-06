package loops

import "fmt"

func Workshop1() {
	aklimdakiSayi := 80
	tahminiSayi := 0
	sayac := 0
	var cond bool = false

	fmt.Print("Bir sayı tahmin ediniz: ")
	fmt.Scanln(&tahminiSayi)
	sayac++

	for tahminiSayi != aklimdakiSayi {
		if tahminiSayi > aklimdakiSayi {
			fmt.Print("Aşağı: ")
			fmt.Scanln(&tahminiSayi)
		} else {
			fmt.Print("Yukarı: ")
			fmt.Scanln(&tahminiSayi)
		}
		if sayac == 9 {
			sayac++
			break
		}
		sayac++
	}

	if tahminiSayi == aklimdakiSayi {
		fmt.Println("Doğru Tahmin")
		cond = true
	}

	if sayac <= 3 {
		fmt.Printf("%v. Tahminde Bildiniz. Süper\n", sayac)
	} else if sayac <= 6 {
		fmt.Printf("%v. Tahminde Bildiniz. Güzel\n", sayac)
	} else if sayac > 6 && cond {
		fmt.Printf("%v. Tahminde Bildiniz. Fena Değil\n", sayac)
	} else if !cond {
		fmt.Println("Tahmin hakkınız bitti", sayac)
	}

	fmt.Println("Oyun Bitti")
}
