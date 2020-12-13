package main

import "fmt"

func main() {
	defer check("first")
	defer check("second")
	check("third")
	defer check("fouth")
	defer check("fifth")
}

func check(s string) {
	fmt.Println(s)
}
