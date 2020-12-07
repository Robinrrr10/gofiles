package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")
	fmt.Println("When is wednesday")
	today := time.Now().Weekday()
	fmt.Println("today is:", today)
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far")
	}
	fmt.Println("end")
}
