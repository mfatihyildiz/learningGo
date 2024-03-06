package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Çift Sayı: ", i)
		time.Sleep(2 * time.Second)
	}
}

func TekSayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Tek Sayı: ", i)
		time.Sleep(1 * time.Second)
	}
}

/* main içi
func main() {
	go goroutines.TekSayilar()
	go goroutines.CiftSayilar()
	time.Sleep(5 * time.Second)
	fmt.Println("Main bitti")
}
*/
