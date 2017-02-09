package main

import "fmt"

type person struct {
	first string
}

func (p person) speak(){
	fmt.Println("Hello", p.first)
}


func main(){
	p1 := person{"Alex"}
	p1.speak()

	xi:=[]int{0,1,2,3,4,5}
	for j:= range xi{
		fmt.Println(j,xi[j])
	}

	m:= map[string]int{"phosykeo":27,"mercy":32,"watson":41}

	for k,v := range m{
		fmt.Println(k,v)
	}

}
