package main

import (
	"fmt"
	"math"
	"strings"
)

type shape interface {
	area() float64
	perimiter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// Methods
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimiter() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimiter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return (4 / 3) * math.Pi * math.Pow(c.radius, 3)
}

func printCircle(c circle) {
	fmt.Println("shape: ", c)
	fmt.Println("area: ", c.area())
	fmt.Println("perimiter: ", c.perimiter())
}

func printRectangle(r rectangle) {
	fmt.Println("shape: ", r)
	fmt.Println("area: ", r.area())
	fmt.Println("perimiter: ", r.perimiter())
}

func print(s shape) {
	fmt.Println("shape: ", s)
	fmt.Println("area: ", s.area())
	fmt.Println("perimiter: ", s.perimiter())
}

func main() {
	rect := rectangle{
		width:  2.0,
		height: 4.12,
	}

	ball := circle{
		radius: 5,
	}

	printRectangle(rect)
	printCircle(ball)
	// refactored to single function print by introducing interface
	fmt.Println(strings.Repeat("=", 64))
	print(rect)
	print(ball)

	var s shape = circle{
		radius: 6.2,
	}

	fmt.Printf("%T\n", s)

	basketball, ok := s.(circle)
	_ = basketball

	if ok {
		fmt.Println("basketball volume: ", basketball.volume())
	}

	// type switch
	s = rectangle{
		width:  1.2,
		height: 3.6,
	}
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has cicle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}
}
