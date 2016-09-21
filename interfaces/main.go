package main

import "fmt"

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 2 * 3.14 * c.radius
}

func info(s Shape) {
	//fmt.Printf("Type %s", s)
	fmt.Printf("Area: %v\n", s.area())
}

func main() {
	s1 := Square{side:10.0}
	info(s1)
	c1 := Circle{radius:5.0}
	info(c1)
}
