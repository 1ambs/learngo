package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(time.Second)
		c <- 42
	}()
	fmt.Println("Wait 1 second")
	fmt.Println(<-c)
	fmt.Println("Channel blocked and unblocked")
}
