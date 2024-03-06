package defer_statement

import "fmt"

type products struct {
	productName string
	unitPrice   int
}

func (p products) save() {
	fmt.Println("ürün eklendi: ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := products{"Laptop", 5000}
	defer p.save()
	p = products{"Mouse", 100}
	fmt.Println("İşlem başarılı")
	fmt.Println(p.productName)
}
