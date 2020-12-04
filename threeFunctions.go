package main

import "fmt"

func add(x int, y int) int {
	var addresult = x + y
	return addresult
}

func sub(x, y int) int { //as x and y are same datatype we can give only one  time  for last
	var addresult = x - y
	return addresult
}

func cal(x int, y int) (int, int) {
	var addresult = x + y
	var subresult = x - y
	return addresult, subresult
}

func cal2(x int, y int) (addresult, subresult int) { //intializing the return datatype here itself
	addresult = x + y
	subresult = x - y
	return //not giving the variable here
}

func main() {
	num1 := 54
	num2 := 12
	var result = add(num1, num2)
	fmt.Println(result)
	subresult := sub(num1, num2)
	fmt.Println(subresult)
	num1 = 5
	num2 = 2
	addr, subr := cal(num1, num2)
	fmt.Println(addr, subr)
	num1 = 9
	num2 = 4
	addr, subr = cal2(num1, num2)
	fmt.Println(addr, subr)
	fmt.Println("Working")
}
