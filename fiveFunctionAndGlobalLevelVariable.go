package main

import "fmt"

var a = 2 //package variable
var b = 5 //package variable

func myfunc() {
	var b = 8  //function variable
	var c = 10 //function variable
	fmt.Println("in my function a ", a)
	fmt.Println("in my function b ", b)
	fmt.Println("in my function c ", c)
}

func main() {
	myfunc()
	fmt.Println("in main function a ", a)
	fmt.Println("in main function b ", b)
}
