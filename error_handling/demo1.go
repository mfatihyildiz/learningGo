package error_handling

import (
	"fmt"
	"os"
)

//type assertion
func Demo1() {
	f, err := os.Open("demo1.txt")

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı: ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
		}
	} else {
		fmt.Println(f.Name())
	}
}

//demo1.txt main2.go ile aynı yerde olmalı demo1.go ile değil

//Interfaces demo3 te devamı mevcut
