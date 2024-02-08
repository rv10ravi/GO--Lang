//Q.6 Develop a scenario based on your domain that incorporates looping, control structures, variables, and constants
package main

import "fmt"

// Constants for rental rates per day
const (
	economyRate    = 30
	standardRate   = 40
	luxuryRate     = 50
	insuranceRate  = 10
)

func main() {
	fmt.Println("Welcome to the Car Rental Management System!")

	for {
		var carType string
		var daysRented int
		var insurance bool

		// Taking user for input
		fmt.Println("\nPlease select the type of car (economy/standard/luxury):")
		fmt.Scanln(&carType)

		fmt.Println("Enter the number of days rented:")
		fmt.Scanln(&daysRented)

		fmt.Println("Do you want insurance coverage? (true/false):")
		fmt.Scanln(&insurance)

		// Calculating rental cost
		var baseRate int
		switch carType {
		case "economy":
			baseRate = economyRate
		case "standard":
			baseRate = standardRate
		case "luxury":
			baseRate = luxuryRate
		default:
			fmt.Println("Invalid car type selected. Please select either economy, standard, or luxury.")
			continue 
		}

		totalCost := baseRate * daysRented

		if insurance {
			totalCost += insuranceRate * daysRented
		}

		// Displaying total rental cost
		fmt.Printf("Total rental cost: %d\n", totalCost)

		// Asking  if the user wants to rent another car
		var anotherCar string
		fmt.Println("Do you want to rent another car? (yes/no)")
		fmt.Scanln(&anotherCar)

		if anotherCar != "yes" {
			break  // Exit the loop
		}
	}

	fmt.Println("Thank you for using the Car Rental Management System!")
}
