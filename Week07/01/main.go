package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/",dog)
	http.HandleFunc("/cat", cat)
	http.HandleFunc("/bug", bug)
	http.HandleFunc("/fish", fish)
	http.HandleFunc("/rat", rat)
	http.HandleFunc("/img/toby.gif", dogPic)
	http.HandleFunc("/img/cat.gif", catPic)
	http.HandleFunc("/img/bug.jpg", bugPic)
	http.HandleFunc("/img/fish.jpg", fishPic)
	http.HandleFunc("/img/rat.jpg", ratPic)
	http.ListenAndServe(":8080",nil)
}

func dog(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="img/toby.gif">`)
}

func cat(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="img/cat.gif">`)
}
func bug(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="img/bug.jpg">`)
}
func fish(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="img/fish.jpg">`)
}
func rat(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="img/rat.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "img/toby.gif")
}
func catPic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "img/cat.gif")
}
func bugPic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "img/bug.jpg")
}
func fishPic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "img/fish.jpg")
}
func ratPic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "img/rat.jpg")
}

