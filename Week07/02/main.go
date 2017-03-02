package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/", dog)
	http.HandleFunc("/cat",cat)
	http.HandleFunc("/bug",bug)
	http.HandleFunc("/fish",fish)
	http.HandleFunc("/rat",rat)
	http.Handle("/resources/",http.StripPrefix("/resources", http.FileServer(http.Dir("./img"))))
	http.ListenAndServe(":8080", nil)
}
func dog(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.gif">`)
}
func cat(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/cat.gif">`)
}
func bug(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/bug.jpg">`)
}
func fish(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/fish.jpg">`)
}
func rat(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/rat.jpg">`)
}
