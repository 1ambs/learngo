package main

import "fmt"

func main() {
	s := "David Marubbi"
	fmt.Println(s, len(s))
	for _, v := range s {
		fmt.Printf("%c", v)
	}
}
