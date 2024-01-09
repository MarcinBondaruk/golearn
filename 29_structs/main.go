package main

import "fmt"

func main() {
	title1, author1, year1 := "The Diving Comedy", "Dante Aligheri", 1320

	type book struct {
		title  string
		author string
		year   int
	}

	book1 := book{
		title:  title1,
		author: author1,
		year:   year1,
	}
	aBook := book{title: "My Biography"}

	fmt.Printf("%T\n %v\n %#v\n", book1, book1, book1)
	fmt.Printf("%#v\n", aBook)

	fmt.Println(book1.title)

	// copy a struct
	book2 := book1
	book2.title = "Another dante book"

	fmt.Printf("%+v\n", book1)
	fmt.Printf("%+v\n", book2)

	// anonymus struct
	diana := struct {
		firstname, lastname string
		age                 int
	}{
		firstname: "diana",
		lastname:  "muller",
		age:       25,
	}

	fmt.Printf("%+v\n", diana)

	// anonymous fields
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{
		"1940 by George Orwell",
		10.2,
		false,
	}

	fmt.Printf("%+v\n", b1)

	type Contact struct {
		email   string
		address string
		phone   int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 115000,
		contactInfo: Contact{
			email:   "johnkeller@gmail.com",
			address: "221b Baker Street",
			phone:   999323121,
		},
	}

	fmt.Printf("%+v\n", john)
	fmt.Printf("Employee's email: %q\n", john.contactInfo.email)
}
