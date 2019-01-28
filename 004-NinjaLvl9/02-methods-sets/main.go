package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
	age       int
}

type dog struct{}

func (r *dog) speak() string {
	return "Woof Woof"
}

func (r person) speak() string {
	return "Hello, my name is " + r.firstname + " " + r.lastname
}

type speaker interface {
	speak() string
}

func saysomething(s speaker) {
	fmt.Println("I have this to say: ", s.speak())
}

func main() {
	a := person{
		firstname: "David",
		lastname:  "Marubbi",
		age:       47,
	}
	fmt.Println(a.speak())

	saysomething(a)
	saysomething(&a)
	var b dog

	saysomething(&b)
}
