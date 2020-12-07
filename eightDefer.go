package main

import "fmt"

func main() {
	a()
	c()
}

func a() {
	fmt.Println("a starts")
	defer b() // defer will run at the  last
	fmt.Println("a end")
}

func b() {
	fmt.Println("BBB")
}

func c() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i) //defer will run at last. If there is more value, then it will work like last in first out
	}
	fmt.Println("bye bye")
}
