//A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type. 
//Declaring a structure:

//  type Address struct {
// 	name string 
// 	street string
// 	city string
// 	state string
// 	Pincode int
// }

 
// package main
 
// import "fmt"
 
// // Defining a struct type
// type Address struct {
//     Name    string
//     city    string
//     Pincode int
// }
 
// func main() {
 
//     // Declaring a variable of a `struct` type
//     // All the struct fields are initialized 
//     // with their zero value
//     var a Address 
//     fmt.Println(a)
 
//     // Declaring and initializing a
//     // struct using a struct literal
//     a1 := Address{"Akshay", "Dehradun", 3623572}
 
//     fmt.Println("Address1: ", a1)
 
//     // Naming fields while 
//     // initializing a struct
//     a2 := Address{Name: "Anikaa", city: "Ballia",
//                                  Pincode: 277001}
 
//     fmt.Println("Address2: ", a2)
 
//     // Uninitialized fields are set to
//     // their corresponding zero-value
//     a3 := Address{Name: "Delhi"}
//     fmt.Println("Address3: ", a3)
// }

//----------------------------------------------

//Pointers in Go programming language or Golang is a variable which is used to store the memory address of another variable. 
// package main
 
// import "fmt"
 
// // defining a structure
// type Employee struct {
//     firstName, lastName string
//     age, salary int
// }
 
// func main() {
 
//     // passing the address of struct variable
//     // emp8 is a pointer to the Employee struct
//     emp8 := &Employee{"Sam", "Anderson", 55, 6000}
 
//     // (*emp8).firstName is the syntax to access
//     // the firstName field of the emp8 struct
//     fmt.Println("First Name:", (*emp8).firstName)
//     fmt.Println("Age:", (*emp8).age)
// }

//----------------------------------------------

// Nested Structure in Golang
// Syntax: 

// type struct_name_1 struct{
//   // Fields
// } 
// type struct_name_2 struct{
//   variable_name  struct_name_1

// }

// package main
 
// import "fmt"
 
// // Creating structure
// type Author struct {
//     name   string
//     branch string
//     year   int
// }
 
// // Creating nested structure
// type HR struct {
 
//     // structure as a field
//     details Author
// }
 
// func main() {
 
//     // Initializing the fields
//     // of the structure
//     result := HR{
     
//         details: Author{"Sona", "ECE", 2013},
//     }
 
//     // Display the values
//     fmt.Println("\nDetails of Author")
//     fmt.Println(result)
// }

//----------------------------------------------
// Anonymous Structure and Field in Golang

// syntax:
// variable_name := struct{
// // fields
// }{// Field_values}

package main 
  
import "fmt"
  
// Main function 
func main() { 
  
    // Creating and initializing 
    // the anonymous structure 
    Element := struct { 
        name      string 
        branch    string 
        language  string 
        Particles int
    }{ 
        name:      "Pikachu", 
        branch:    "ECE", 
        language:  "C++", 
        Particles: 498, 
    } 
  
    // Display the anonymous structure 
    fmt.Println(Element) 
} 