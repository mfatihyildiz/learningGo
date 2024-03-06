package slices

import "fmt"

func Demo2() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)

	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)

	sehirler = append(sehirler, "Antalya")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:3])

}
