package variables

import "fmt"

func Demo1() {

	var metin string = "Merhaba"
	fmt.Println(metin)

	var kdv int = 45
	fmt.Println(kdv)

	var kdv2 float32 = 62.8
	fmt.Println(kdv2)

	kdv3 := 83.62
	fmt.Println(kdv3)
	fmt.Printf("Veri tipi: %T\n", kdv3)

	var durum bool
	var metin1 string = "Hasan"
	var metin2 string = "Hüseyin"
	durum = metin1 == metin2
	fmt.Println(durum)
	durum = metin1 != metin2
	fmt.Println(durum)

	fmt.Println(!durum)

}
