package main

import "fmt"

func main() {
	var num []int = []int{56, 23, 90, 11, 35, 12, 33, 50, 7, 16}
	fmt.Println(num)
	var s1 []int = num[2:5]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	var p1 []int //default length  and capacity is 0
	fmt.Println("length:", len(p1))
	fmt.Println("capacity:", cap(p1))

	var p2 []int = make([]int, 3)
	fmt.Println(p2)
	fmt.Println("length:", len(p2))
	fmt.Println("capacity:", cap(p2))

	var p3 []int = make([]int, 3, 10)
	fmt.Println(p3)
	fmt.Println("length:", len(p3))
	fmt.Println("capacity:", cap(p3))

}
