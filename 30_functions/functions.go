package main

import (
	"fmt"
	"math"
	"strings"
)

type Person struct {
	fullName string
	age      int
}

func fun1() {
	fmt.Println("This is called from func 1")
}

func fun2(a int, b int) {
	fmt.Println("Sum: ", a+b)
}

func f3(a, b, c int, d string, s float64) {
	fmt.Println(1, b, c, d, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

// named return parameter
func sum(a, b int) (s int) {
	fmt.Println(s)

	s = a + b
	// naked return
	return
}

// variadic functions - take variadic number of arguments 0+
// only last parameter of a function can be variadic
// when numer of params is unknown - use variadic func

func myVariadic1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func NewPerson(age int, names ...string) Person {
	return Person{
		age:      age,
		fullName: strings.Join(names, " "),
	}
}

func foo() {
	fmt.Println("This is Foo")
}

func bar() {
	fmt.Println("This is bar")
}

func foobar() {
	fmt.Println("This is FooBar")
}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	fun1()
	fun2(5, 7)
	f3(1, 2, 3, "golang", 4.2)
	fmt.Println(f4(6.4))
	a, b := f5(5, 7)
	fmt.Println(a, b)
	mySum := sum(6, 5)
	fmt.Println(mySum)

	myVariadic1(1, 2, 3, 4, 5)
	myVariadic1()

	nums := []int{3, 4, 5, 6, 7}
	nums = append(nums, 3, 5, 6)
	myVariadic1(nums...)
	sarah := NewPerson(33, "Sarah", "Jessica", "Parker")
	fmt.Printf("%+v,\n", sarah)
	fmt.Printf("%#v\n", sarah)

	for _, v := range sarah.fullName {
		fmt.Println(string(v))
	}

	foo()
	bar()

	//because of defer, foo will be called after main function exits, resulting in being called after bar
	// defer is useful for cleaning after function closing file for example
	defer foo()
	bar()

	fmt.Println("Just a string after defering foo() and calling bar()")
	// defer calls are put on a stack so it executs from latest to oldest defered func
	defer foobar()

	// anonymous functions
	func(msg string) {
		fmt.Println(msg)
	}("Hello from anonymous function")

	msgPrinter := func(msg string) {
		fmt.Println(msg)
	}

	msgPrinter("Im an anonymous function assigned to variable")

	incrementer := increment(10)
	fmt.Printf("%T\n", incrementer)
	incrementer()
	fmt.Println(incrementer())

}
