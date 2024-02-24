package main

import "fmt"

func main(){
	var sum float64 = 0
	var temp[7] float64
	var avg float64
	fmt.Println("Weclcome to tem records")
	fmt.Println("Enter Your 7 days Temperature")
	for i:=0;i<7;i++{
		fmt.Scanln(&temp[i])
		sum += temp[i]
	}
	avg = sum/7
	fmt.Println("Average ",avg)

	fmt.Println("Days Where temp exceeds 30 degree")
	for i:=0;i<7;i++{
		if temp[i]>30.0 {
			fmt.Println("Day:",i+1, " Temp:",temp[i])
		}
	}


}