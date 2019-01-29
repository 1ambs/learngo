package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println("Wait 1 second")
	fmt.Println(<-c)

	fmt.Println("Channel blocked and unblocked")
}
