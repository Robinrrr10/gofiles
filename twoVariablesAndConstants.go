package main

import "fmt"

func main() {
	var num = 2
	fmt.Println(num)
	var num1 = 4
	var num2 = 3
	var result = num1 + num2
	fmt.Println(result)
	var num3 int = 3
	var num4 int = 1
	num4 = 5
	var res2 int
	res2 = num3 + num4
	fmt.Println(res2)
	var num5, num6 int
	num5, num6 = 4, 7
	fmt.Println(num5 + num6)
	num7 := 19
	fmt.Println(num7)
	var num9 int = num2 + num3
	fmt.Println(num9)

	const num8 = 121 //constant variable
	fmt.Println(num8)
}
