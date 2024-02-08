// Q1. You're developing an online store application in GoLang. As part of the application, you need to keep track of various product details such as name, price, and quantity in stock. Design a set of variables and assign values to represent a specific product in the inventory. Ensure you use appropriate data types for each variable to accurately capture the information.
package main

import "fmt"

//using struct to store the product details
type product struct {
	name string
	price float64
	quantity int
}


func main() {
	//Declaring arrray of structures
	var p[100] product

	var n int

	//Taking the inputs from user
	fmt.Println("Enter the number of products you want to enter:")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++ {
		fmt.Println("Enter Product ",i," Name:")
		fmt.Scanln(&p[i].name)
		fmt.Println("Enter Product",i," Price:")
		fmt.Scanln(&p[i].price)
		fmt.Println("Enter Product ",i," Quantity:")
		fmt.Scanln(&p[i].quantity)
		fmt.Println("Product ",i," details entered successfully")
		fmt.Println("====================================")
	}
	for i:=1;i<=n;i++ {
		//Printing the product details
		fmt.Println("Product ",i," Name:",p[i].name)
		fmt.Println("Product ",i,"Price:",p[i].price)
		fmt.Println("Product ",i," Quantity:",p[i].quantity)
		fmt.Println("====================================")
	}
}