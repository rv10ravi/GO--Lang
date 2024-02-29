//You are given a list of integers. Your task is to find all the pairs of integers in the list whose sum is 
//equal to a given target value. However, each integer in the list can only be used once in a pair. 
//Write a Go function findPairs that takes in a list of integers numbers, and an integer target, 
//and returns a slice of pairs of integers that sum up to the target. Array elements and the target should be given by the user

//Suppose the array elements are 2, 7, 11, 15, 3, 6, 8, 12    and the target value is  18, the expected output is 
//Pairs with sum 18 are: [[7 11] [6 12]]

package main

import "fmt"

func main(){
	var numbers = []int{2, 7, 11, 5, 4, 6, 8, 12}
	var target int
	fmt.Println("Enter the target value")
	fmt.Scanln(&target)
	fmt.Println("Pairs with sum",target,"are:",findPairs(numbers,target))
}

func findPairs(numbers []int, target int) [][]int {
	var pairs [][]int
	for i:=0;i<len(numbers);i++ {
		for j:=i+1;j<len(numbers);j++ {
			if numbers[i]+numbers[j] == target {
				pairs = append(pairs,[]int{numbers[i],numbers[j]})
			}
		}
	}
	return pairs
}