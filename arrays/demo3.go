package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{60, 30, 50, 80, 2}
	fmt.Println(sayilar)

	enBuyuk := 0

	for i := 0; i < len(sayilar); i++ {
		fmt.Println(sayilar[i])
	}

	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}
	}

	fmt.Printf("En büyük sayı: %v", enBuyuk)
}
