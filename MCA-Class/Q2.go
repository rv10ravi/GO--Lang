// Q2 You're tasked with building a student information system in GoLang for a school. Each student record needs to store details such as student ID, name, age, and grade. Define variables to store the information of a single student and assign values accordingly. Pay attention to selecting appropriate data types to represent each piece of information.
package main

import "fmt"

//using struct to store the student details
type student struct {
	id int
	name string
	age int
	grade string
}

func main(){
	//Declaring arrray of structures
	var s[100] student
	var n int

	//Taking the inputs from user
	fmt.Println("Enter the number of students you want to enter:")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++ {
		fmt.Println("Enter Student ",i," ID:")
		fmt.Scanln(&s[i].id)
		fmt.Println("Enter Student ",i," Name:")
		fmt.Scanln(&s[i].name)
		fmt.Println("Enter Student ",i," Age:")
		fmt.Scanln(&s[i].age)
		fmt.Println("Enter Student ",i," Grade:")
		fmt.Scanln(&s[i].grade)
		fmt.Println("Student details entered successfully")
		fmt.Println("====================================")
	}
	for i:=1;i<=n;i++ {
		//Printing the product details
		fmt.Println("Student ",i," ID:",s[i].id)
		fmt.Println("Student ",i," Name:",s[i].name)
		fmt.Println("Student ",i," Age:",s[i].age)
		fmt.Println("Student ",i," Grade:",s[i].grade)
		fmt.Println("====================================")
	}

}