package functions

import "fmt"

func Topla(sayi1 int, sayi2 int) int {
	toplam := sayi1 + sayi2
	return toplam
}

func SelamVer(isim string) {
	fmt.Print("Merhaba " + isim + "\n")
}

/* main2.go
functions.SelamVer("Hasan")
functions.SelamVer("Hüseyin")

var sonuc = functions.Topla(3, 5)
fmt.Println(sonuc * 10)
*/
