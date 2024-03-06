package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	rate               float64
	adress             string
}

type Car struct {
	creditPaymentTotal float64
	rate               float64
	carInfo            string
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlypayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{10, 100000, "Ankara"}
	credit2 := Mortgage{12, 50000, "İstanbul"}
	credit3 := Car{15, 60000, "Volvo"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlypayment(credits)

	fmt.Println("Toplam ödeme: ", total)
}
