package main

import "fmt"

type person struct{
	name string
}
func (a person) speak (){
	fmt.Println("hello",a.name)
}

func main(){
	p1:= person{"alex"}
	p1.speak()

	xi:=[]int{0,1,2,3,4,5}
	for j:= range xi{
		fmt.Println(j,xi[j])
	}

	m:= map[string]int{"phosykeo":27,"bond":40,"Gates":62}

	for k,v := range m{
		fmt.Println(k,v)
	}
}