//Q4 You are tasked with creating a grade calculator program in GoLang. The program should prompt the user to input their scores in three subjects: Math, Science, and English. Based on these scores, calculate the average grade (assuming each subject is equally weighted) and display the corresponding letter grade (A, B, C, D, or F) using control flow.

package main

import "fmt"

func main() {
	//Declaring  variables
	var math, science, english float64

	//Taking the inputs from user
	fmt.Println("Enter your marks in Math:")
	fmt.Scanln(&math)
	fmt.Println("Enter your marks in Science:")
	fmt.Scanln(&science)
	fmt.Println("Enter your marks in English:")
	fmt.Scanln(&english)

	//Calculating the average grade
	avg := (math+science+english)/3

	//Displaying the corresponding letter grade
	if avg>=90 {
		fmt.Println("Your Grade is A")
	} else if avg>=80 && avg<90 {
		fmt.Println("Your Grade is B")
	} else if avg>=70 && avg<80 {
		fmt.Println("Your Grade is C")
	} else if avg>=60 && avg<70 {
		fmt.Println("Your Grade is D")
	} else {
		fmt.Println("Your Grade is F")
	}
}