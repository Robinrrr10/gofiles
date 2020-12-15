package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Mpage struct {
	Tit     string
	BodyMes string
	Para    string
}

func main() {
	fmt.Println("Start")
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("mydynpage.html")
		data := "hiiiiiiiiiiiiiiiiiii"
		t.Execute(res, data)
	})

	http.HandleFunc("/check", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("mydynpage2.html")
		pa := Mpage{Tit: "My titles learn", BodyMes: "This  is my heeeeeeee", Para: "Paraaaaaaaaaaaaaaaaaaaaa"}
		t.Execute(res, pa)
	})
	http.ListenAndServe(":8301", nil)
}
