package main

import (
	"fmt"
	"time"
)

//it is like multi thread concept
func main() {
	fmt.Println("start")
	go logTheValue("first") //it wont wait till the function. It is async
	go logTheValue("two")
	go logTheValue("three")
	fmt.Println("end")
	time.Sleep(time.Second)
	fmt.Println("end2")
}

func logTheValue(s string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(s, ":", i)
	}
}
