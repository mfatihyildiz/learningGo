package conditionals

import "fmt"

func Workshop1() {
	var deger1, deger2 int = 80, 40
	deger3 := 600

	if deger1 > deger2 {
		if deger1 > deger3 {
			fmt.Printf("En Büyük Değer: %v", deger1)
		} else {
			fmt.Printf("En Büyük Değer: %v", deger3)
		}
	} else if deger2 > deger3 {
		fmt.Printf("En Büyük Değer: %v", deger2)
	} else {
		fmt.Printf("En Büyük Değer: %v", deger3)
	}
	/*
		var sayi1, sayi2, sayi3 int = 10, 00, 30
		var enBuyuk int = sayi1

		if sayi2 > enBuyuk {
			enBuyuk = sayi2
		}

		if sayi3 > enBuyuk {
			enBuyuk = sayi3
		}
		fmt.Printf("En Büyük Değer: %v", enBuyuk)
	*/
}

func Workshop2() {
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			fmt.Printf("%v Çift sayı\n", i)
		} else {
			fmt.Printf("%v Tek sayı\n", i)
		}
	}
}
