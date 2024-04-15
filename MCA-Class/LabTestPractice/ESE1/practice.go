package main

import (
	"errors"
	"fmt"
)


type LibraryItem interface {
    Borrow() error
    Return() error
    Available() bool
}


type Book struct {
    Title       string
    Author      string
    IsAvailable bool
}


func NewBook(title, author string) *Book {
    return &Book{
        Title:       title,
        Author:      author,
        IsAvailable: true,
    }
}


func (b *Book) Borrow() error {
    if !b.IsAvailable {
        return errors.New("Book is not available for borrowing")
    }
    b.IsAvailable = false
    return nil
}


func (b *Book) Return() error {
    if b.IsAvailable {
        return errors.New("Book is already available")
    }
    b.IsAvailable = true
    return nil
}


func (b *Book) Available() bool {
    return b.IsAvailable
}


func AddToLibrary(items ...LibraryItem) {
    for _, item := range items {
        if book, ok := item.(*Book); ok {
            library = append(library, book)
        } else {
            fmt.Println("Error: Attempted to add non-book item to the library")
        }
    }
}

var library []*Book

func main() {
    
    book1 := NewBook("Atom", "Nuclear")
    book2 := NewBook("Happiness", "Life")

    AddToLibrary(book1, book2)

    var choice int
    for {
        fmt.Println("Library Menu:")
        fmt.Println("1. Borrow a Book")
        fmt.Println("2. Return a Book")
        fmt.Println("3. Check Book Availability")
        fmt.Println("4. Exit")
        fmt.Print("Enter your choice: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            borrowBook()
        case 2:
            returnBook()
        case 3:
            checkAvailability()
        case 4:
            fmt.Println("Exiting program.")
            return
        default:
            fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
        }
    }
}

func borrowBook() {
    var title string
    fmt.Print("Enter the title of the book you want to borrow: ")
    fmt.Scanln(&title)

    for _, book := range library {
        if book.Title == title {
            if err := book.Borrow(); err != nil {
                fmt.Println(err)
                return
            }
            fmt.Printf("You have borrowed the book '%s'\n", title)
            return
        }
    }
    fmt.Printf("Book '%s' not found in the library\n", title)
}

func returnBook() {
    var title string
    fmt.Print("Enter the title of the book you want to return: ")
    fmt.Scanln(&title)

    for _, book := range library {
        if book.Title == title {
            if err := book.Return(); err != nil {
                fmt.Println(err)
                return
            }
            fmt.Printf("You have returned the book '%s'\n", title)
            return
        }
    }
    fmt.Printf("Book '%s' not found in the library\n", title)
}

func checkAvailability() {
    var title string
    fmt.Print("Enter the title of the book you want to check: ")
    fmt.Scanln(&title)

    for _, book := range library {
        if book.Title == title {
            if book.Available() {
                fmt.Printf("The book '%s' is available\n", title)
            } else {
                fmt.Printf("The book '%s' is not available\n", title)
            }
            return
        }
    }
    fmt.Printf("Book '%s' not found in the library\n", title)
}
