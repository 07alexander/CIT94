package main

import "fmt"

type person struct{
	fName string
	lName string
	favFood []string
}
func main(){
	p1 := person{"alex", "phosykeo",
		[]string{"spaghetti", "pho", "tacos" },
	}
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)

	for i:= range p1.favFood{
		fmt.Println(p1.favFood[i])
	}

}
