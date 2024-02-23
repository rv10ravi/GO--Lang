//You are tasked with creating an inventory management system for a small grocery store. Design a Go program that utilizes a struct to represent products with attributes such as name, price, and quantity. Implement a slice to store instances of these product structs. Allow the user to perform operations such as adding new products, updating existing ones, and removing products from the inventory. Utilize looping and control structures to provide a user-friendly interface.

package main

import "fmt"

type Product struct{
	Name string
	Price float64
	Quantity int

}

var inventory []Product

func main(){
	for {
		fmt.Println("\n--- Welcome To Grocery Store Management System ---")
		fmt.Println("1. Add a Product")
		fmt.Println("2. Remove a Product")
		fmt.Println("3. Update Product Details")
		fmt.Println("4. List All Products")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addProduct()
		case 2:
			removeProduct()
		case 3:
			updateProduct()
		case 4:
			listProducts()
		case 5:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addProduct() {
	var name string
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	// Check if the product already exists
	for _, product := range inventory {
		if product.Name == name {
			fmt.Println("Product already exists in the inventory")
			return
		}
	}

	var price float64
	var quantity int

	fmt.Print("Enter Price: ")
	fmt.Scanln(&price)
	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&quantity)

	inventory = append(inventory, Product{Name: name, Price: price, Quantity: quantity})
}

func removeProduct() {
	var name string
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	for i, product := range inventory {
		if product.Name == name {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Println("Product removed from the inventory")
			return
		}
	}

	fmt.Println("Product not found in the inventory")
}

func updateProduct() {
	var name string
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	for i, product := range inventory {
		if product.Name == name {
			var price float64
			var quantity int

			fmt.Print("Enter Price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter Quantity: ")
			fmt.Scanln(&quantity)

			inventory[i].Price = price
			inventory[i].Quantity = quantity
			fmt.Println("Product details updated")
			return
		}
	}

	fmt.Println("Product not found in the inventory")
}

func listProducts() {
	fmt.Println("\n--- List of Products ---")
	for _, product := range inventory {
		fmt.Println("--------------------")
		fmt.Println("Name:", product.Name)
		fmt.Println("Price:", product.Price)
		fmt.Println("Quantity:", product.Quantity)
		fmt.Println("--------------------")
	}
}

