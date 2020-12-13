package main

import "fmt"

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 { //when we use same method name which is in interface, then it will automatically implement from that interface
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	height, width float64
}

func (r Rectangle) area() float64 {
	return 3.14 * r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {

	fmt.Println("start")
	c := Circle{radius: 10}
	r := Rectangle{height: 2, width: 4}
	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}
