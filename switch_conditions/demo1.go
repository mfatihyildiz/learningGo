package switch_conditions

import "fmt"

func Demo1() {
	grade := -3

	switch grade {
	//switch grade := -3; grade {
	case 5:
		fmt.Println("Pek İyi")
	case 4:
		fmt.Println("İyi")
	case 3:
		fmt.Println("Orta")
	case 2:
		fmt.Println("Geçer")
	case 1:
		fmt.Println("Başarısız")
	default:
		fmt.Println("Geçersiz Not")
	}
}
