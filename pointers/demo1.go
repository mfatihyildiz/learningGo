package pointers

import "fmt"

func Demo1(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Demodaki sayı:", *sayi)

}

/* main içi
func main() {
	sayi := 20
	pointers.Demo1(&sayi)
	fmt.Println("Maindeki sayı:", sayi)
}
*/
