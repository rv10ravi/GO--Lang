//Q5 You want to build a simple interest calculator in GoLang. The program should ask the user to input multiple sets of data, each containing the principal amount, the annual interest rate, and the number of years for which the interest is to be calculated. Use arrays to store the input data for each set, calculate the simple interest for each set using the formula: Simple Interest = (Principal Amount * Annual Interest Rate * Number of Years) / 100, and display the results.

package main

import "fmt"

type simpleInterest struct {
	principalAmount float64
	annualInterestRate float64
	numberOfYears int
}

func main() {

	//Declaring arrray of structures
	var s[100] simpleInterest
	var n int

	//Taking the inputs from user
	fmt.Println("Enter the number of sets of data you want to enter:")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++ {
		fmt.Println("Enter Set ",i," Principal Amount:")
		fmt.Scanln(&s[i].principalAmount)
		fmt.Println("Enter Set ",i," Annual Interest Rate:")
		fmt.Scanln(&s[i].annualInterestRate)
		fmt.Println("Enter Set ",i," Number of Years:")
		fmt.Scanln(&s[i].numberOfYears)
		fmt.Println("Set ",i," details entered successfully")
		fmt.Println("====================================")
	}
	for i:=1;i<=n;i++ {
		//Calculating the simple interest
		si := (s[i].principalAmount * s[i].annualInterestRate * float64(s[i].numberOfYears)) / 100
		//Displaying the simple interest
		fmt.Println("Set ",i," Simple Interest:",si)
		fmt.Println("====================================")
	}
	
}
