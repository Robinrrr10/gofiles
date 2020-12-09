package main

import "fmt"

func main() {
	myarr()
	myarr2()
}

func myarr() {
	var num = [8]int{18, 23, 12, 89, 43, 6, 91, 28}
	for i := 0; i < 8; i++ {
		fmt.Printf("value of index %d is: %d \n", i, num[i])
	}
}

func myarr2() {
	var num []int
	num = []int{8, 123, 42, 89, 93, 61, 14, 48}
	for i, x := range num {
		fmt.Println(i, x)
	}
}
