package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
}

type Author struct {
	Name    string
	Email   string
	Address Address
}

type Book struct {
	ID          int
	Title       string
	ISBN        string
	Author      Author
	PublishDate string
	Genre       string
	Price       float64
	Currency    string
}

type Bookstore struct {
	Books []Book
}

func main() {
	// Simulated JSON data with multiple books
	jsonData := []byte(`{
		"Books": [
			{
				"ID": 1,
				"Title": "The God of Small Things",
				"ISBN": "9780679457312",
				"Author": {
					"Name": "Arundhati Roy",
					"Email": "arundhati@gmail.com",
					"Address": {"City": "New Delhi", "State": "Delhi", "Country": "India"}
				},
				"PublishDate": "April 4, 1997",
				"Genre": "Fiction",
				"Price": 899,
				"Currency": "Rupees"
			},
			{
				"ID": 2,
				"Title": "The Namesake",
				"ISBN": "9780618485222",
				"Author": {
					"Name": "Jhumpa Lahiri",
					"Email": "jhumpa@gmail.com",
					"Address": {"City": "Kolkata", "State": "West Bengal", "Country": "India"}
				},
				"PublishDate": "September 1, 2003",
				"Genre": "Fiction",
				"Price": 950,
				"Currency": "Rupees"
			}
		]
	}`)

	var bookstore Bookstore

	// Unmarshalling
	if err := json.Unmarshal(jsonData, &bookstore); err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	// Print unmarshalled data
	fmt.Println("Unmarshalled data:", bookstore.Books)
}
