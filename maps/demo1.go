package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["Table"] = "Masa"
	sozluk["Book"] = "Kitap"
	sozluk["Pencil"] = "Kalem"

	fmt.Println(sozluk["Book"])
	fmt.Println("Eleman sayısı:", len(sozluk))

	delete(sozluk, "Book")
	fmt.Println("Eleman sayısı:", len(sozluk))
	fmt.Println("Sözlük:", sozluk)

	deger, varMi := sozluk["Table"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu:", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)
}
