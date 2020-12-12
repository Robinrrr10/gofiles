package main

import "fmt"

func main() {
	var num int = 12
	fmt.Println("Value of num is:", num)
	fmt.Println("Address of num is:", &num)

	var ptr1 *int
	ptr1 = &num
	fmt.Println("Value of ptr1 is:", ptr1)
	fmt.Println("Address of ptr1 is:", &ptr1)
	fmt.Println("Value by giving address ptr1 is:", *ptr1)

	var ptr2 **int
	ptr2 = &ptr1
	fmt.Println("Value of ptr2 is:", ptr2)
	fmt.Println("Address of ptr2 is:", &ptr2)
	fmt.Println("Value by giving address ptr2 is:", **ptr2)
}
