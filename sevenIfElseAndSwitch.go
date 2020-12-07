package main

import "fmt"

func main() {
	num := 4
	fmt.Println("given number  is:", num)
	if num <= 5 {
		fmt.Println("number is less than or equal to 5")
	} else if num > 5 {
		fmt.Println("number is greater than 5")
	}
	printNum(num)
}

func printNum(n int) {
	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		{
			fmt.Println("two")
		}
	case 3:
		{
			fmt.Println("three")
		}
	case 4:
		fmt.Println("four")
	default:
		fmt.Println("Big number")
	}
}
