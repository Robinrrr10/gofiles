package main

import "fmt"

func main() {
	a()
}
func a() {
	fmt.Println("a start")
	b()
	fmt.Println("a end")
}

func b() {
	fmt.Println("b")
}
