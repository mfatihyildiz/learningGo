package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)
	fmt.Println(isimler)

	isimler[0] = "Hasan"
	isimler[1] = "Hüseyin"
	isimler[2] = "Rıfkı"
	fmt.Println(isimler)

	isimler = append(isimler, "Nuri")
	fmt.Println(isimler)
}
