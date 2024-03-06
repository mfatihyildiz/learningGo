package loops

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya"

	i := 0
	for i < 5 {
		fmt.Println(metin)
		i++
	}
	fmt.Println("Bitti")
}
