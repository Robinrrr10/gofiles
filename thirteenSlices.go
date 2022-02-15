package main

import "fmt"

func main() {

	myslice()

	myslice2()
}

func myslice() {
	var num = [8]int{3, 81, 51, 2, 33, 21, 22, 42}
	var sl1 []int = num[2:5] // index 2 to length 5 of the num
	fmt.Println(sl1)
	var sl2 []int = num[:5] //default 0th index to length 5 of the num
	fmt.Println(sl2)
	var sl3 []int = num[2:] //index 2 to upto end of the length of the num
	fmt.Println(sl3)
}

func myslice2() {
	var num = [8]int{3, 81, 51, 2, 33, 21, 22, 42}
	var sl1 []int = num[2:5]
	fmt.Println(sl1)
	num[3] = 66 //When changing value in array, it will reflect in slice which uses that array
	fmt.Println(sl1)
}
