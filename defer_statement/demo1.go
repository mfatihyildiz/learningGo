package defer_statement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}

func C() {
	fmt.Println("C fonksiyonu çalıştı")
}

func B() {
	defer A()
	defer C()
	fmt.Println("B fonksiyonu çalıştı")
}
