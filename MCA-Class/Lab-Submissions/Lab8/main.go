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
	Author      Author
	Publisher   string
	ISBN        string
	PublishDate string
	Genre       string
	Price       float64
	Currency    string
	Tags        []string
	Location    string
}

type Bookstore struct {
	ID       int
	Name     string
	Location string
	Books    []Book
}

// UnmarshalBookstore unmarshals JSON data into a Bookstore struct
func UnmarshalBookstore(data []byte) (Bookstore, error) {
	var bookstore Bookstore
	err := json.Unmarshal(data, &bookstore)
	return bookstore, err
}

// MarshalBookstore marshals a Bookstore struct into JSON data
func MarshalBookstore(bookstore Bookstore) ([]byte, error) {
	return json.Marshal(bookstore)
}

func main() {
	jsonData := []byte(`{
		"ID": 1,
		"Name": "Indian Book Emporium",
		"Location": "456 Gandhi Road, New Delhi",
		"Books": [
			{
				"ID": 1,
				"Title": "The God of Small Things",
				"Author": {
					"Name": "Arundhati Roy",
					"Email": "arundhati@gmail.com",
					"Address": {"City": "New Delhi", "State": "Delhi", "Country": "India"}
				},
				"Publisher": "Random House",
				"ISBN": "9780679457312",
				"PublishDate": "April 4, 1997",
				"Genre": "Fiction",
				"Price": 499,
				"Currency": "Rupees",
				"Tags": ["novel", "contemporary"],
				"Location": "Fiction Section"
			},
			{
				"ID": 2,
				"Title": "The Namesake",
				"Author": {
					"Name": "Jhumpa Lahiri",
					"Email": "jhumpa@gmail.com",
					"Address": {"City": "Kolkata", "State": "West Bengal", "Country": "India"}
				},
				"Publisher": "Houghton Mifflin",
				"ISBN": "9780618485222",
				"PublishDate": "September 1, 2003",
				"Genre": "Fiction",
				"Price": 560,
				"Currency": "Rupees",
				"Tags": ["novel", "diaspora"],
				"Location": "Fiction Section"
			}
		]
	}`)

	bookstore, err := UnmarshalBookstore(jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	fmt.Println("Unmarshalled data:", bookstore.Books)

	bookstoreData, err := MarshalBookstore(bookstore)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println(" \n \n Marshalled data:", string(bookstoreData))
}
