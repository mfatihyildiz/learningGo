package error_handling

import "fmt"

type BorderException struct {
	parameter int
	message   string
}

func (b *BorderException) Error() string {
	return fmt.Sprintf("%d-----%s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &BorderException{tahmin, "Sınırların dışında kaldı"}
	}
	return "Bildiniz", nil
}

/*main içi
fmt.Println(error_handling.TahminEt2(120))
*/
