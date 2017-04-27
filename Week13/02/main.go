package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
}

func main(){
	p1:=person{
		First:"Alexander",
		Last:"Phosykeo",
	}

	xp:= []person{p1}

	fmt.Printf("go data: %+v\n", xp)

	bs, err:= json.Marshal(xp)
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println("json", string(bs))
}
