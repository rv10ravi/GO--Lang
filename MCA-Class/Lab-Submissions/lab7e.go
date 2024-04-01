package main

import "fmt"

type Book struct {
    ID          int
    Title       string
    Author      string
    Price       float64
    Stock       int
}

func displayBook(book Book) {
    fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f , Stock: %d \n", book.ID, book.Title, book.Author, book.Price , book.Stock)
}

func updateStock(book *Book, quantity int) {
    book.Stock += quantity
}

func main() {
    book := Book{1, "Atomic", "James", 150, 10}
	displayBook(book)
    updateStock(&book, 5) 
    fmt.Println("Updated Stock:", book.Stock) 
    displayBook(book)
}

