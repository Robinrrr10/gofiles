package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Start")
	res, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	sttt, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(sttt))
	fmt.Println("end")
}
