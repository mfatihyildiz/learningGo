package pointers

import "fmt"

func Demo2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki sayı:", sayilar[0])

}

/* main içi
func main() {
	sayilar := []int{1, 2, 3}
	pointers.Demo2(sayilar)
	fmt.Println("Maindeki sayı:", sayilar[0])

}
*/
