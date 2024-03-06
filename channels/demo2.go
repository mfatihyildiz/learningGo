package channels

import "fmt"

func merhaba(chan1 chan string) {
	chan1 <- "Merhaba"
}

func Demo2() {
	myChannel := make(chan string)

	go merhaba(myChannel)

	fmt.Println(<-myChannel)
}
