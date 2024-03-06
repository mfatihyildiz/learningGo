package switch_conditions

import "fmt"

func Demo2() {
	switch x := 25; {
	case x < 20:
		fmt.Printf("%d küçüktür 20\n", x)
		fallthrough
	case x < 50:
		fmt.Printf("%d küçüktür 50\n", x)
		fallthrough

	case x < 100:
		fmt.Printf("%d küçüktür 100\n", x)
		fallthrough

	case x < 200:
		fmt.Printf("%d küçüktür 200\n", x)
	}
}
