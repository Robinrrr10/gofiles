package main

import "fmt"

func infi() {
	for {
		fmt.Println("this is infinite loop")
	}
}

func main() {
	//infi()
	myForLoop()
	myForLoopInSingleLine()
}

func myForLoop() {
	i := 1
	for i <= 5 {
		fmt.Println("my loop da", i)
		i = i + 1 // or i++
	}
}

func myForLoopInSingleLine() {
	for i := 1; i <= 5; i++ {
		fmt.Println("the loop", i)
	}
}
