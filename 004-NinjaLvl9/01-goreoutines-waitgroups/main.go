package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(7)

	go func() {
		fmt.Println("hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing two")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing three")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing four")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing five")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing six")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing seven")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}