package slices

import "fmt"

func Demo3() {
	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)

	mySlc := underArray[2:8]
	fmt.Println(mySlc)

	mySlc2 := underArray[:6]
	fmt.Println(mySlc2)

	mySlc3 := underArray[3:]
	fmt.Println(mySlc3)

	mySlc4 := underArray[:]
	fmt.Println(mySlc4)

	mySlc[0] = 100
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	fmt.Println(mySlc3)
	fmt.Println(mySlc4)

	//
	fmt.Println("///////////////////")
	//

	mySlc5 := append(mySlc3, 4, 5)
	fmt.Println(mySlc5)

	mySlc3[0] = 200
	fmt.Println(mySlc3)
	fmt.Println(mySlc5)

	//
	fmt.Println("///////////////////")
	//

	mySlc6 := []int{1, 2, 3}
	mySlc7 := []int{4, 5}

	mySlc6 = append(mySlc6, mySlc7...)
	fmt.Println(mySlc6)

	//
	fmt.Println("///////////////////")
	//

	fmt.Println(mySlc)

	mySlc = mySlc[2:]
	//ilk n elemanı silmek [n:]
	fmt.Println(mySlc)

	mySlc = mySlc[:len(mySlc)-3]
	//son n elemanı silmek [:len(mySlc)-n]
	fmt.Println(mySlc)

}
