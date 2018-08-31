package main

import "fmt"

type Book struct {
	Name string
	page int
}

func passByValue(book Book) {
	book.Name = "Hello"
}

func passByPoint(book *Book) {
	book.Name = "Hello"
}

func main() {
	book := Book{
		"test", 100,
	}

	passByValue(book)
	fmt.Println(book)

	passByPoint(&book)
	fmt.Println(book)
}
