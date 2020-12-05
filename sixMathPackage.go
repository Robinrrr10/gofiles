package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64 = 12
	var res1 = math.Sqrt(num1)
	fmt.Println("result is:", res1)
	fmt.Printf("result is %f", res1)
	fmt.Println()
	fmt.Printf("result is %g", res1)
	fmt.Println()
	fmt.Printf("result is %.2f", res1)
	fmt.Println()
	fmt.Printf("result is %.3f", res1)
	fmt.Println()
	var res2 = math.Round(res1)
	fmt.Printf("result is %f", res2)
}
