// This is the documentation for a simple program
package main

import (
	"fmt"
)

// Dave is a function to output a message from dave.
func dave() {
	fmt.Println("Dave says Hello")
}

func main() {
	// some more text
	func() {
		fmt.Println("Hello go world")
	}()
}
