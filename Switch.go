// switch optstatement; optexpression{
// 	case expression1: Statement..
// 	case expression2: Statement..
// 	...
// 	default: Statement..
// 	}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Switch Statement")
	// switch statement
	// switch statement is used to compare the value of an expression to a list of values and execute the corresponding block of code.
	// The switch statement is used to replace multiple if else statements.

	// fmt.Println("Enter a number between 1 to 5")
	// var num int
	// fmt.Scan(&num)

	// switch num {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// case 5:
	// 	fmt.Println("Five")
	// default:
	// 	fmt.Println("Invalid Number")
	// }


	//-----------------Switch without an expression-----------------
	var value int = 2 
      
    // Switch statement without an      
    // optional statement and  
    // expression 
   switch { 
       case value == 1: 
       fmt.Println("Hello") 
       case value == 2: 
       fmt.Println("Bonjour") 
       case value == 3: 
       fmt.Println("Namstay") 
       default:  
       fmt.Println("Invalid") 
   } 
}
