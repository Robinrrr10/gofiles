package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello. This is simple text")
	})

	http.HandleFunc("/man2", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Contect-Type", "html/text")
		fmt.Fprintf(res, "<html><h3>My heading</h3></html>")
	})
	http.HandleFunc("/color", myColorFullpage)
	http.ListenAndServe(":8200", nil)
}

func myColorFullpage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "myhtpage.html")
}
