package main

import "fmt"

func main(){
	for{
		fmt.Println("Welcome to the vending machine!")
		fmt.Println("1. Coke")
		fmt.Println("2. Pepsi")
		fmt.Println("3. Water")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			CalculateAmount(choice)
		case 2:
			CalculateAmount(choice)
		case 3:
			fmt.Println("Water dispensed")
			fmt.Println("Water is Free , Enjoy")
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")

	}

}
	

}

func CalculateAmount(choice int) {
	cokePrice := 20
	pepsiPrice := 25

	var amount int
	println("Enter the amount")
	fmt.Scanln(&amount)

	if choice == 1 {
		if amount >= cokePrice {
			fmt.Println("Coke dispensed")
			fmt.Println("Change:", amount-cokePrice)
		} else {
			fmt.Println("Insufficient amount")
		}
	} else if choice == 2 {
		if amount >= pepsiPrice {
			fmt.Println("Pepsi dispensed")
			fmt.Println("Change:", amount-pepsiPrice)
		} else {
			fmt.Println("Insufficient amount")
		}
	}


}