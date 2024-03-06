package error_handling

import (
	"errors"
	"math"
)

func Demo4(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negatif sayıların karekökü alınamaz")
	}
	return math.Sqrt(num), nil
}

/* main içi
result, err := error_handling.Demo4(-40)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result)
}
*/
