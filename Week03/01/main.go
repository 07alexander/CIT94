//Initialize a SLICE of int using a composite literal; print out the slice; range over the slice printing out just the index; range over the slice printing out both the index and the value
package main

import "fmt"

func main(){
	xi:=[]int{0,1,2,3,4,5}

	for i:= range xi{
		fmt.Println(i)
	}
	for j:= range xi{
		fmt.Println(j,xi[j])
	}

}
