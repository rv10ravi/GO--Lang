//Q.Utilize the principles of Variables, Control Flow, Arrays, Slices, Maps, and Structs
// within a specified domain of your choice. Begin the program with a comprehensive overview detailing 
//the scenario and the concepts being implemented. Ensure adequate comments are provided throughout the code.
// Evaluation of this program will be based on the following criteria:
//R1: Concept Clarity/Viva: 8 Marks
//R2: Correctness: 8 Marks
//R3: Validations( only with if statement expected): 8 Marks
//R4: Ability to Relate to Real-Time Scenario: 8 Marks
//R5: Complexity: 8 Marks

// My domain is bookstore management
// The program is a simple bookstore management system that allows the user to add, remove, update, list, and search for books in the inventory

package main

import (
	"fmt"
)

// Defining a struct to represent a book
type Book struct {
	Title    string
	Author   string
	Price    float64
	Quantity int
}

// Global variable to store book inventory
var inventory map[string]Book


// array to store the names of books that not available
var notAvailableBooks []string


//Main function
func main() {
	// Initializing the inventory map
	inventory = make(map[string]Book)

	// Adding some initial books to the inventory
	inventory["2347140"] = Book{"Mahabharat", "Vyasa", 1450, 100}
	inventory["2344589"] = Book{"Ramayan", "Valmiki", 2375, 150}

	// Main menu 
	for {
		fmt.Println("\n--- Welcome To Bookstore Management System ---")
		fmt.Println("1. Add a Book")
		fmt.Println("2. Remove a Book")
		fmt.Println("3. Update Book Details")
		fmt.Println("4. List All Books")
		fmt.Println("5. Search Book by Title")
		fmt.Println("6. Print Not Available Books")
		fmt.Println("7. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBook()
		case 2:
			removeBook()
		case 3:
			updateBook()
		case 4:
			listBooks()
		case 5:
			searchBook()
		case 6:
			printNotAvailableBooks()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

// Function to add a book to the inventory
func addBook() {
	var isbn string
	fmt.Print("Enter ISBN: ")
	fmt.Scanln(&isbn)

	// Check if the book already exists (Validation For I)
	if _, ok := inventory[isbn]; ok {
		fmt.Println("Book already exists in the inventory")
		return
	}

	var title, author string
	var price float64
	var quantity int

	fmt.Print("Enter Title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter Author: ")
	fmt.Scanln(&author)
	fmt.Print("Enter Price: ")
	fmt.Scanln(&price)
	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&quantity)

	inventory[isbn] = Book{title, author, price, quantity}
	fmt.Println("Book added to inventory")
}

// Function to remove a book from the inventory
func removeBook() {
	var isbn string
	fmt.Print("Enter ISBN: ")
	fmt.Scanln(&isbn)

	// Check if the book exists
	if _, ok := inventory[isbn]; !ok {
		fmt.Println("Book does not exist in the inventory")
		return
	}

	delete(inventory, isbn)
	fmt.Println("Book removed from inventory")
}

// Function to update the details of a book in the inventory
func updateBook() {
	var isbn string
	fmt.Print("Enter ISBN: ")
	fmt.Scanln(&isbn)

	// Check if the book exists
	if _, ok := inventory[isbn]; !ok {
		fmt.Println("Book does not exist in the inventory")
		return
	}

	var title, author string
	var price float64
	var quantity int

	fmt.Print("Enter Title: ")
	fmt.Scanln(&title)
	fmt.Print("Enter Author: ")
	fmt.Scanln(&author)
	fmt.Print("Enter Price: ")
	fmt.Scanln(&price)
	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&quantity)

	inventory[isbn] = Book{title, author, price, quantity}
	fmt.Println("Book details updated")
}

// Function to list all books in the inventory
func listBooks() {
	fmt.Println("\n--- List of Books ---")
	for isbn, book := range inventory {
		fmt.Println("--------------------")
		fmt.Println("ISBN:", isbn)
		fmt.Println("Title:", book.Title)
		fmt.Println("Author:", book.Author)
		fmt.Println("Price:", book.Price)
		fmt.Println("Quantity:", book.Quantity)
		fmt.Println("--------------------")
	}
}

// Function to search for a book by title
func searchBook() {
	var title string
	fmt.Print("Enter Title to search: ")
	fmt.Scanln(&title)

	found := false
	for isbn, book := range inventory {
		if book.Title == title {
			fmt.Println("--------------------")
			fmt.Println("ISBN:", isbn)
			fmt.Println("Title:", book.Title)
			fmt.Println("Author:", book.Author)
			fmt.Println("Price:", book.Price)
			fmt.Println("Quantity:", book.Quantity)
			fmt.Println("--------------------")
			found = true
		}
	}

	if !found {
		fmt.Println("Book not found in the inventory")
		notAvailableBooks = append(notAvailableBooks, title)
	}
}

// Function to print the names of books that are not available
func printNotAvailableBooks() {
	fmt.Println("\n--- Books Not Available ---")
	for _, title := range notAvailableBooks {
		fmt.Println(title)
	}

	//Printing only first two books using slices 
	fmt.Println("\n--- First Two Books Not Available ---")
	for i,title := range notAvailableBooks[:2] {
		fmt.Println(i+1,"-",title)
	}
}