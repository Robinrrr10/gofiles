package main

import "fmt"

func main() {
	x := 3

	switch {
	case x == 1:
		fmt.Println("one")
	case x == 2:
		fmt.Println("two")
	case x == 3:
		fmt.Println("three")
	default:
		fmt.Println("more than 3")
	}

}
