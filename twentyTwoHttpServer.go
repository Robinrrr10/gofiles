package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start")
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hi hellow")
	})

	http.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hi all users")
	})

	myHandle := func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hi friends")
	}
	http.HandleFunc("/frd", myHandle)
	http.HandleFunc("/grt", greatUsrs)
	http.ListenAndServe(":8100", nil)

	fmt.Println("Stop")
}

func greatUsrs(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hi greater users")
}
