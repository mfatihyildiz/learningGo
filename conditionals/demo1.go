package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki Para Yetersiz")
	}

	if hesap <= cekilmekIstenen {
		fmt.Println("Para hazırlanıyor")
	}

	hesap -= cekilmekIstenen
	//fmt.Println("Hesapta Kalan Para: " + fmt.Sprintf("%v", hesap))
	fmt.Printf("Hesapta kalan Para: %v", hesap)
}
