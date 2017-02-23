package main

import (
	"html/template"
	"net/http"
)

type person struct{
	first string
	age int
}

var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	letters:=[]string{"a","b","c","d"}
	grade:=map[string]int{
		"A":4,
		"B":3,
		"C":2,
		"D":1,
		"F":0,

	}
	p1:= person{
		"Alex",
		27,
	}
	tpl.ExecuteTemplate(w,"index.gohtml", 10)
	tpl.ExecuteTemplate(w,"index.gohtml", "a string")
	tpl.ExecuteTemplate(w,"index.gohtml", true)
	tpl.ExecuteTemplate(w,"index.gohtml", letters)
	for k := range grade{
		tpl.ExecuteTemplate(w,"index.gohtml", k)

	}
	for _,v := range grade{
		tpl.ExecuteTemplate(w,"index.gohtml", v)

	}
	tpl.ExecuteTemplate(w,"index.gohtml", p1)

}


