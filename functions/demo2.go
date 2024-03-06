package functions

func Dortislem(sayi1 int, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2)
	return toplam, cikarim, carpim, bolum
}

/* main2.go
sonuc1, sonuc2, sonuc3, sonuc4 := functions.Dortislem(10, 6)
//sonuc1, sonuc2, sonuc3, _ := functions.Dortislem(10, 6)

fmt.Println("Toplam: ", sonuc1)
fmt.Println("Çıkarım: ", sonuc2)
fmt.Println("Çarpım: ", sonuc3)
fmt.Println("Bölüm: ", sonuc4)
*/
