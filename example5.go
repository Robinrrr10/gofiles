package main

import "fmt"

func main() {
	num1 := 19
	num2 := 12
	a, b := cal(num1, num2)
	fmt.Println(a, b)
}

var cal = func(x, y int) (int, int) { //writing function differently
	return x + y, x - y
}
