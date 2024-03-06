package structs

import "fmt"

func Demo3() {
	type employee struct {
		name      string
		age       int
		isMarried bool
	}

	type manager struct {
		employee
		hasDegree bool
	}

	e1 := employee{"Hasan", 40, true}
	fmt.Println(e1)

	m1 := manager{
		employee:  employee{"Ayşe", 28, false},
		hasDegree: true,
	}
	fmt.Println(m1)

}
