package main

import "fmt"

func main() {

	var s1 Student = Student{12, "Ram", 56}
	fmt.Println(s1)
	fmt.Println(s1.name)

	var s2 Student = Student{rollNo: 12, mark: 67}
	fmt.Println(s2)
	fmt.Println(s2.rollNo)
	fmt.Println(s2.name)

}

type Student struct {
	rollNo int
	name   string
	mark   int
}
