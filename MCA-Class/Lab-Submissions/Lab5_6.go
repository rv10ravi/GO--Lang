package main

import (
	"errors"
	"fmt"
)

type Book struct {
	ISBN   string
	Title  string
	Author string
	Price  float64
}

type BookStore interface {
	addBook(book Book) error
	findBookByISBN(isbn string) (Book, error)
	buyBook(isbn string) error
	generateBill() float64
}

type BookStoreImpl struct {
	books map[string]Book
	bill  float64
}

func newBookStore() *BookStoreImpl {
	return &BookStoreImpl{
		books: make(map[string]Book),
		bill:  0,
	}
}

func (bs *BookStoreImpl) addBook(book Book) error {
	if _, exists := bs.books[book.ISBN]; exists {
		return errors.New("Book with this ISBN already exists")
	}
	bs.books[book.ISBN] = book
	return nil
}

func (bs *BookStoreImpl) findBookByISBN(isbn string) (Book, error) {
	book, exists := bs.books[isbn]
	if !exists {
		return Book{}, errors.New("Book not found")
	}
	return book, nil
}

func (bs *BookStoreImpl) buyBook(isbn string) error {
	book, exists := bs.books[isbn]
	if !exists {
		return errors.New("Book not found")
	}
	bs.bill += book.Price
	delete(bs.books, isbn)
	return nil
}

func (bs *BookStoreImpl) generateBill() float64 {
	return bs.bill
}

func main() {
	bs := newBookStore()

	for {
		fmt.Println("Book Store Management System")
		fmt.Println("Menu:")
		fmt.Println("1. Add a book")
		fmt.Println("2. Find a book by ISBN")
		fmt.Println("3. Buy a book")
		fmt.Println("4. Generate Bill")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var isbn, title, author string
			var price float64

			fmt.Print("Enter ISBN: ")
			fmt.Scanln(&isbn)
			fmt.Print("Enter Title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter Author: ")
			fmt.Scanln(&author)
			fmt.Print("Enter Price: ")
			fmt.Scanln(&price)

			err := bs.addBook(Book{ISBN: isbn, Title: title, Author: author, Price: price})
			if err != nil {
				fmt.Println("Error:", err.Error())
			} else {
				fmt.Println("Book added successfully!")
			}
			fmt.Println("-----------------------------------")

		case 2:
			var searchISBN string
			fmt.Print("Enter ISBN to search: ")
			fmt.Scanln(&searchISBN)

			foundBook, err := bs.findBookByISBN(searchISBN)
			if err != nil {
				fmt.Println("Error:", err.Error())
			} else {
				fmt.Println("Book found:")
				fmt.Printf("Title: %s\nAuthor: %s\nPrice: %.2f\n", foundBook.Title, foundBook.Author, foundBook.Price)
			}
			fmt.Println("-----------------------------------")
		case 3:
			var buyISBN string
			fmt.Print("Enter ISBN to buy: ")
			fmt.Scanln(&buyISBN)

			err := bs.buyBook(buyISBN)
			if err != nil {
				fmt.Println("Error:", err.Error())
			} else {
				fmt.Println("Book purchased successfully!")
			}
			fmt.Println("-----------------------------------")

		case 4:
			totalBill := bs.generateBill()
			fmt.Printf("Total Bill: %.2f\n", totalBill)
			fmt.Println("-----------------------------------")

		case 5:
			fmt.Println("Exiting...")
			
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
	}
}
}