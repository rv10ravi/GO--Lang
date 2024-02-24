package main

import "fmt"

func main(){
	var choice int 
	for{
		fmt.Println("Welcome To Vending Machine")
		fmt.Println("1: Water ")
		fmt.Println("2: Coke ")
		fmt.Println("3: Pepsi ")
		fmt.Println("----------------- ")
		fmt.Println("Enter Your Choice:")
		fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Water is Free")
		fmt.Println("Enjoy the drink")
		
	case 2:
		Purchase(choice)
	case 3:
		Purchase(choice)
	}
	}
}

func Purchase(choice int) {
	pepsiPrice := 20
	cokePrice := 15
	var full_amount int
	fmt.Println("Enter The AMount Your are Giving:")
	fmt.Scanln(&full_amount)

	if choice == 2 {
		return_amount := full_amount - cokePrice
		fmt.Println("REturn amount is ", return_amount)
	} else {
		return_amount := full_amount - pepsiPrice
		fmt.Println("REturn amount is ", return_amount)
	}
}