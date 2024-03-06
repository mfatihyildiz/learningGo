package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Added: ", c.firstName)
}
func (c customer) update() {
	fmt.Println("Updated: ", c.firstName)
}

func Demo2() {
	c := customer{"Hasan", "Hüseyin", 35}
	c.save()
	c.update()
}
