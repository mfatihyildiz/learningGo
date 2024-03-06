package structs

import "fmt"

type family struct {
	name      string
	age       int
	isMarried bool
}

func showFamily() []family {
	family1 := family{"Hasan", 25, true}
	family2 := family{"Hüseyin", 5, false}
	family3 := family{"Ayşe", 61, true}

	return []family{family1, family2, family3}
}

func Workshop2() {

	families := showFamily()

	for i := 0; i < len(families); i++ {
		fmt.Printf("%v, %T\n", families[i], families[i])
	}
}
