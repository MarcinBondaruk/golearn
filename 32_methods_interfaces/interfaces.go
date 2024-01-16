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

	circle := circle{
		radius: 5,
	}

	printRectangle(rect)
	printCircle(circle)
	// refactored to single function print by introducing interface
	fmt.Println(strings.Repeat("=", 64))
	print(rect)
	print(circle)

}
