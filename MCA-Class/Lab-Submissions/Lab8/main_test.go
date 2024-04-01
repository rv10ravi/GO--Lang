package main

import (
	"testing"
)

func TestUnmarshalBookstore(t *testing.T) {
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

	_, err := UnmarshalBookstore(jsonData)
	if err != nil {
		t.Errorf("Error unmarshalling: %v", err)
	}
}
