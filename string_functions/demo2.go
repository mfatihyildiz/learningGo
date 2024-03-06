package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Hasan"

	fmt.Println(s.HasPrefix(isim, "Has"))
	fmt.Println(s.HasSuffix(isim, "an"))
	fmt.Println(s.Index(isim, "sa"))

	harfler := []string{"h", "a", "s", "a", "n"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	//sondaki int değeri kaç tanesinin
	//değişeceğini gösterir
	//Hepsinin değişmesi için -1
	fmt.Println(s.Replace(sonuc, "*", "+", 1))

	fmt.Println(s.Split(sonuc, "*"))
	fmt.Println(s.Split(sonuc, "-"))
	fmt.Println(s.Repeat(sonuc, 5))
}
