package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/m", monday)
	http.HandleFunc("/t", tuesday)
	http.HandleFunc("/w", wednesday)
	http.HandleFunc("/th", thursday)
	http.HandleFunc("/f", friday)
	http.HandleFunc("/s", saturday)
	http.HandleFunc("/su", sunday)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"index.gohtml", nil)
}

func monday(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"monday.gohtml", nil)
}

func tuesday(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"tuesday.gohtml", nil)
}
func wednesday(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"wednesday.gohtml", nil)
}
func thursday(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"thursday.gohtml", nil)
}
func friday(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"friday.gohtml", nil)
}
func saturday(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"saturday.gohtml", nil)
}

func sunday(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w,"sunday.gohtml", nil)
}

