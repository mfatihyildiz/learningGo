package string_functions

import (
	"fmt"
	s "strings" //alias
)

func Demo1() {
	isim := "Hasan"
	//case sensitive
	fmt.Println(s.Count(isim, "h"))
	fmt.Println(s.Contains(isim, "g"))
	fmt.Println(s.Index(isim, "g"))
	fmt.Println(s.Index(isim, "s"))

	sonuc := s.Index(isim, "a")
	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))

}
