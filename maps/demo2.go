package maps

import "fmt"

func Demo2() {
	myMap := map[string]int{
		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 70,
	}

	fmt.Println(myMap)
	fmt.Println(myMap["Ahmet"], myMap["Selim"])

	sg := myMap //maps----pass by reference
	fmt.Println(sg)

	delete(sg, "Ahmet")
	fmt.Println(sg)
	fmt.Println(myMap)

	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
