package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID          int
	Title       string
	Author      string
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

func main() {
	bookstore := Bookstore{
		ID:       1,
		Name:     "Indian Book Emporium",
		Location: "456 Gandhi Road, New Delhi",
		Books: []Book{
			{
				ID:          1,
				Title:       "The God of Small Things",
				Author:      "Arundhati Roy",
				Publisher:   "Random House",
				ISBN:        "9780679457312",
				PublishDate: "April 4, 1997",
				Genre:       "Fiction",
				Price:       499,
				Currency:    "Rupees",
				Tags:        []string{"novel", "contemporary"},
				Location:    "Fiction Section",
			},
			{
				ID:          2,
				Title:       "The Namesake",
				Author:      "Jhumpa Lahiri",
				Publisher:   "Houghton Mifflin",
				ISBN:        "9780618485222",
				PublishDate: "September 1, 2003",
				Genre:       "Fiction",
				Price:       560,
				Currency:    "Rupees",
				Tags:        []string{"novel", "diaspora"},
				Location:    "Fiction Section",
			},
		},
	}

	bookstoreData, err := json.Marshal(bookstore)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println("Marshalled data:", string(bookstoreData))
}
