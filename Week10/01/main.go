package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
)

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session")
	if err != nil{
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session",
			Value: id.String(),
			HttpOnly: true,
			Path:"/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)

}
