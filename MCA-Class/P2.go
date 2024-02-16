package main
import "fmt"

func main(){
	var sum float64 = 0
	var temp[7] float64
	var avg float64

	fmt.Println("Enter the 7 temperatures")
	for i:=0;i<7;i++ {
		fmt.Scanln(&temp[i])
		sum += temp[i]
	}
	avg = sum/7
	fmt.Println("Average temperature is",avg)

	fmt.Println("Days where temp exceeds 30 degree")
	for i:=0;i<7;i++ {

		if temp[i] > 30 {
			fmt.Println("Day ",i+1, ":",temp[i])
		}
	}


}

