package main

import (
	"fmt"
	"time"
)

type names []string

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

// this will not modify c values
func (c car) changeCar(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCarWithPointerReciever(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// n is a method reciever
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)
	fmt.Printf("seconds in a day %v\n", seconds)

	friends := names{"Dan", "Marry"}
	friends.print()

	myCar := car{
		brand: "Peugot",
		price: 56000,
	}

	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)
	myCar.changeCar("Renault", 20000)
	fmt.Println(myCar)
	myCar.changeCarWithPointerReciever("Ferrari", 1000000)
	fmt.Println(myCar)
}
