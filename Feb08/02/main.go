package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request){
	fmt.Println(w,"Hello from foo")
}