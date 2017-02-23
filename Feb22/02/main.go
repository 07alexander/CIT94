package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/",dog)
	http.HandleFunc("/toby.gif", dogPic)
	http.ListenAndServe(":8080",nil)
}

func dog(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="toby.gif">`)
}

func dogPic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "toby.gif")
}
