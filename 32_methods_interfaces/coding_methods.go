package main

import "fmt"

type money float64
type book struct {
	price float64
	title string
}

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func (b book) vat() float64 {
	return b.price * 1.09
}

func (b *book) discount() {
	b.price *= 0.9
}

func main() {
	var change money = 5.887
	var gazylions money = 199234124
	change.print()
	fmt.Println(gazylions.printStr())

	witcher := book{
		title: "Witcher",
		price: 24.99,
	}

	fmt.Println(witcher.vat())
	witcher.discount()
	fmt.Println(witcher.vat())
}
