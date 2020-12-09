package main

import "fmt"

func main() {
	var num int = 1
	for num <= 10 {
		fmt.Println(num)
		if num == 5 {
			break
		}
		num++
	}
}
