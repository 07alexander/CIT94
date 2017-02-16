package main

import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,`<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Title</title>
	</head>
	<body>
	Hello from index
	<br>
	<a href="/">index</a>
	<a href="/about">about</a>
	<a href="/contact">contact</a>
	</body>
	</html>`)
}
func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,`<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Title</title>
	</head>
	<body>
	Hello from about
	<br>
	<a href="/">index</a>
	<a href="/about">about</a>
	<a href="/contact">contact</a>
	</body>
	</html>`)
}
func contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,`<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Title</title>
	</head>
	<body>
	Hello from contact
	<br>
	<a href="/">index</a>
	<a href="/about">about</a>
	<a href="/contact">contact</a>
	</body>
	</html>`)
}

