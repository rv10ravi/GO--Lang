package main

import (
	"fmt"
)

type Book struct {
	title  string
	author string
	price  float64
}

func updatePriceByValue(book Book, newPrice float64)  {
	book.price = newPrice
}

func updatePriceByReference(book *Book, newPrice float64) {
	book.price = newPrice
}

func displayBookInfo(book Book) {
	fmt.Println("Book Title:", book.title)
	fmt.Println("Book Author:", book.author)
	fmt.Printf("Book Price: %.2f\n", book.price)
}

func main() {
	var title, author string
	var price float64

	fmt.Println("Enter book title:")
	fmt.Scanln(&title)

	fmt.Println("Enter book author:")
	fmt.Scanln(&author)

	fmt.Println("Enter book price:")
	fmt.Scanln(&price)

	book := Book{title, author, price}

	fmt.Println("Original Book Info:")
	displayBookInfo(book)

	var newPrice float64

	fmt.Println("Enter new price (call by value):")
	fmt.Scanln(&newPrice)
	updatePriceByValue(book, newPrice)
	fmt.Println("After call by value:")
	displayBookInfo(book)

	fmt.Println("Enter new price (call by reference):")
	fmt.Scanln(&newPrice)
	updatePriceByReference(&book, newPrice)
	fmt.Println("After call by reference:")
	displayBookInfo(book)
}
