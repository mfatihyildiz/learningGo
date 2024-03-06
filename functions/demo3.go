package functions

func ToplaVariadic(sayilar ...int) int {
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam += sayilar[i]
	}
	return toplam
}

/* main2.go
var sonuc = functions.ToplaVariadic(4, 8, 10, 41, 78, 145)
fmt.Println(sonuc)

sayilar := []int{4, 89, 10, 23}
fmt.Println(functions.ToplaVariadic(sayilar...))
*/
