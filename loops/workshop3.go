package loops

import "fmt"

func Workshop3() {
	sayi1, sayi2, arkadas1, arkadas2 := 0, 0, 0, 0

	fmt.Print("İki sayı giriniz:")
	fmt.Scanln(&sayi1, &sayi2)

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			arkadas1 += i
		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			arkadas2 += i
		}
	}

	if arkadas1 == sayi2 && arkadas2 == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayıdır", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayı değildir", sayi1, sayi2)
	}
}
