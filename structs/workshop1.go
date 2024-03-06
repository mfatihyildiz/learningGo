package structs

import "fmt"

type rectangle struct {
	length int
	width  int
}

func (r rectangle) display() {
	fmt.Println("Uzun kenar:", r.length)
	fmt.Println("Kısa kenar:", r.width)
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (r rectangle) circumference() int {
	return 2 * (r.length + r.width)
}

func Workshop1() {
	r := rectangle{10, 7}
	r.display()
	fmt.Println("Alan:", r.area())
	fmt.Println("Çevre:", r.circumference())
}
