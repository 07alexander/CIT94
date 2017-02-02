//Initialize a MAP using a composite literal where the key is a string and the value is an int; print out the map; range over the map printing out just the key; range over the map printing out both the key and the value

package main

import "fmt"

func main(){
	m:=map[string]int{"a":1,"b":2,"c":3,"d":4,"f":5}

	for k:=range m{
		fmt.Println(k)
	}
	for l,v:=range m{
		fmt.Println(l,v)
	}
}
