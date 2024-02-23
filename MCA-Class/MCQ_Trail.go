//Q1
// package main
// import "fmt"
// func main() {
// var x int = 10
// var y float64 = 5.5
// result := x + y
// fmt.Println("Result:", result)
// }

//output: invalid operation: x + y (mismatched types int and float64)

//-----------------------------------------------------------------------------------------------

//Q.2
// package main
// import "fmt"
// func main() {
// var x int = 10
// var y float64 = 5.5
// result := float64(x) + y
// fmt.Println("Result:", result)
// }

//output: Result: 15.5

//-----------------------------------------------------------------------------------------------

//Q.3
// package main
// import "fmt"
// func main() {
// for i := 0; i < 5; i++
// fmt.Println(i)
// }

//output: syntax error: unexpected ++, expecting { after for clause

//-----------------------------------------------------------------------------------------------

// Question: 4
// package main
// import "fmt"
// func main() {

// for i := 0; i< 5; i++ {
// fmt.Println(i)
// }
// }

//output: 0 1 2 3 4

//-----------------------------------------------------------------------------------------------

//Q5
// package main
// import "fmt"
// func main() {
// var x int
// for x < 5 {
// fmt.Println(x)
// }
// }

//output: 000000 in loop

//-----------------------------------------------------------------------------------------------

//Q6
// package main
// import "fmt"
// func main() {
// var x int
// for x < 5 {
// fmt.Println(x)
// x++
// }
// }

//output: 0 1 2 3 4

//-----------------------------------------------------------------------------------------------

// Question: 7
// package main
// import "fmt"
// func main() {
// var i int
// for i := 0; i < 3; i++ {

// fmt.Println(i)
// }
// fmt.Println(i)
// }

//output: 0 1 2 0

//-----------------------------------------------------------------------------------------------

//Q8
// package main
// import "fmt"
// func main() {
// var num int = 5
// if num % 2 == 0 {
// fmt.Println("Even")
// } else
// 	fmt.Println("Odd")

// }

//output: syntax error: unexpected fmt.Println, expecting {

//-----------------------------------------------------------------------------------------------


// Question: 10
// package main
// import "fmt"
// func main() {
// var num int = 5
// if num % 2 == 0 {

// fmt.Println("Even")
// } else {
// fmt.Println("Odd")
// }
// }

//output: Odd

//-----------------------------------------------------------------------------------------------

//Q11 &12
// package main
// import "fmt"
// func main() {
// var i int
// for i := 0; i<3; i++ {
// fmt.Println(i)
// }
// }

//output: 0 1 2

//-----------------------------------------------------------------------------------------------

//Q13
// package main
// import "fmt"
// func main() {
// const x int = 5
// x = 10
// fmt.Println(x)
// }

//output: cannot assign to x

//-----------------------------------------------------------------------------------------------

//Q14
// package main
// import "fmt"
// func main() {
// var x int = 5
// x = 10
// fmt.Println(x)
// }

//output: 10

//-----------------------------------------------------------------------------------------------


//Q15
// package main
// import "fmt"
// func main() {
// var num int = 5
// if num < 10
// fmt.Println("Less than 10")
// else {
// fmt.Println("Greater than or equal to 10")
// }
// }

//output: syntax error: unexpected fmt.Println, expecting {

//-----------------------------------------------------------------------------------------------

//Q16
// package main
// import "fmt"
// func main() {
// var num int = 5
// if num < 10 {
// fmt.Println("Less than 10")
// } else {
// fmt.Println("Greater than or equal to 10")
// }
// }

//output: Less than 10

//-----------------------------------------------------------------------------------------------


//Q17
// package main
// import "fmt"
// func main() {
// for i := 0; i < 3; i++ {
// if i % 2 == 0 {
// fmt.Println("Even")
// } else {
// fmt.Println("Odd")
// }
// }
// }

//output: Even Odd Even

//-----------------------------------------------------------------------------------------------

package main
import "fmt"
func main() {
    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
}
