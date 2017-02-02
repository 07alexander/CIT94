package main

import (
	"fmt"
)

type person struct {
	last string
	age  int
}

func main() {
	p1 := person{"Moneypenny", 24}
	fmt.Println(p1)
	p1.foo()
	p1.walk()
}

func (p person) foo() {
	fmt.Println("Hello from foo")
}

func (a person) walk(){
	fmt.Println(a.last, "is walking")
}

// func (receiver) identifier(parameters) (returns) {}
