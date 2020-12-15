package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name     string
	Password string
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "myform.html")
	})

	http.HandleFunc("/myform", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.ParseFiles("mysucc.html")
		name := req.FormValue("name")
		pass := req.FormValue("password")
		usr := User{Name: name, Password: pass}
		t.Execute(res, usr)
	})

	http.ListenAndServe(":8331", nil)
}
