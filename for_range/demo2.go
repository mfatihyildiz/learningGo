package for_range

import "fmt"

func Demo2() {
	sozluk := map[string]string{"Book": "Kitap", "Table": "Masa"}

	for k /*key*/, v /*value*/ := range sozluk {
		fmt.Println(k)
		fmt.Println(v)
	}
}
