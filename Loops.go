// for initialization; condition; post{
// 	// statements....
// }

// The for loop is a control flow statement that allows code to be executed repeatedly based on a given condition. 
//The for loop in Go is similar to the for loop in other programming languages, but it has a few unique features.

package main

import "fmt"

func main() {

	// for i:=0; i<5; i++{
	// 	fmt.Println(i)
	// }
	//output: 0 1 2 3 4

	// i:=5
	// for i>0{
	// 	fmt.Println(i)
	// }
	//output: 5 5 5 5 IN LOOP

	//Simple range in for loop: You can also use the range in the for loop. Syntax:
	//for i, j:= range rvariable{
   		// statement..
	//}
	
	// arr :=[]int{1,2,3,4,5}
	// for i,j:= range arr{
	// 	fmt.Println("Index:",i," -- Value:",j)
	// }


	//For Maps: A for loop can iterate over the key and value pairs of the map. Syntax:
	// for key, value := range map { 
	//      // Statement.. 
	// }



	// m := map[string]int{
	// 	"one": 1,
	// 	"two": 2,
	// 	"three": 3,
	// 	"four": 4,
	// 	"five": 5,
	// }

	// fmt.Println("Map:",m)
	// for key, value := range m {
	// 	fmt.Println("Key:", key, "Value:", value)
	// }

	//output:
	// Map: map[one:1 two:2 three:3 four:4 five:5]
	// Key: one Value: 1
	// Key: two Value: 2
	// Key: three Value: 3
	// Key: four Value: 4
	// Key: five Value: 5


	



}
