//Keywords - 25
// break
// case
// chan
// const
// continue
// default
// defer
// else
// fallthrough
// for
// func
// go
// goto
// if
// import
// interface
// map
// package
// range
// return
// select
// struct
// switch
// type
// var

//---------------------------------------------
//Basic Data Types:

// Numbers -- Integers , float , complex
// Booleans
// Strings


// ----------------------------------------------
// Declaring Variables 
// var variable_name type = expression
// variable_name:= expression


//-----------------------------------------------
// Operators

package main

import "fmt"

func main(){
	 var a int =10
	 var b int =20

	// Arithemetic operators
	fmt.Println("Addition of a and b is ",a+b)

	fmt.Println("Subtraction of a and b is ",a-b)

	fmt.Println("Multiplication of a and b is ",a*b)

	fmt.Println("Division of a and b is ",a/b)

	fmt.Println("Modulus of a and b is ",a%b)

	//Relational Operators
	fmt.Println("Is a equal to b ",a==b)

	fmt.Println("Is a not equal to b ",a!=b)

	fmt.Println("Is a greater than b ",a>b)

	fmt.Println("Is a less than b ",a<b)

	fmt.Println("Is a greater than or equal to b ",a>=b)

	fmt.Println("Is a less than or equal to b ",a<=b)

	//Logical Operators

	fmt.Println("Logical AND of a and b ",a>b && a<b)

	fmt.Println("Logical OR of a and b ",a>b || a<b)

	fmt.Println("Logical NOT of a ",!(a>b))

	//Bitwise Operators

	fmt.Println("Bitwise AND of a and b ",a&b)

	fmt.Println("Bitwise OR of a and b ",a|b)

	fmt.Println("Bitwise XOR of a and b ",a^b)

	fmt.Println("Bitwise NOT of a ",^a)

	fmt.Println("Bitwise Left Shift of a and b ",a<<b)

	fmt.Println("Bitwise Right Shift of a and b ",a>>b)




	
}