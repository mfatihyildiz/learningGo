package channels

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam += i
	}
	ciftSayiCn <- toplam
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam += i
	}
	tekSayiCn <- toplam
}

/* main içi
func main() {

	ciftSayiToplamCn := make(chan int)
	tekSayiToplamCn := make(chan int)

	go channels.TekSayilar(ciftSayiToplamCn)
	go channels.CiftSayilar(tekSayiToplamCn)

	ciftSayiToplam, tekSayiToplam := <-ciftSayiToplamCn, <-tekSayiToplamCn
	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println("Çarpım: ", carpim)
}
*/
