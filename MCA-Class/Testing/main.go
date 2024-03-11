package main

import "fmt"

// Add function adds two integers and returns the result
func Add(a, b int) int {
    return a + b
}

func main() {
    result := Add(-1, 2)
    fmt.Println("Result of addition:", result)
}