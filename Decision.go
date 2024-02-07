//if , if else , nested-if , if else-if


//program to take input age  and print eligible if it is greater than or equal to 18

package main

import "fmt"

func main(){
	var age int
	fmt.Println("Enter Your Age : ")
	fmt.Scan(&age)
	
	//if statement
	// if age >= 18 {
	// 	fmt.Println("You are eligible ! ")
	// }

	//if else statement
	// if age>= 18 {
	// 	fmt.Println("You are eligible")
	// } else {
	// 	fmt.Println("You are not eligible")
	// }

	//nested if statement 
	// if age>=18{
	// 	if age<25{
	// 		fmt.Println("You are Young")
	// 	}
	// }

	
	//if..else..if ladder
	
	if age>50{
		fmt.Println("You are old")
	} else if age<=18 {
		fmt.Println("You are under age")
	} else {
		fmt.Println("You are eligible")
	}
		

}