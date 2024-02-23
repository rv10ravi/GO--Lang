//Implement the concepts of Interfaces using basic calculator

package main

import "fmt"

type calculator interface {
	add(a int, b int) int
	sub(a int, b int) int
	mul(a int, b int) int
	div(a int, b int) int
}

type calc struct {
}

func (c calc) add(a int, b int) int {
	return a + b
}

func (c calc) sub(a int, b int) int {
	return a - b
}

func (c calc) mul(a int, b int) int {
	return a * b
}

func (c calc) div(a int, b int) int {
	return a / b
}

func main() {
	var c calculator
	c = calc{}
	fmt.Println(c.add(10, 20))
	fmt.Println(c.sub(10, 20))
	fmt.Println(c.mul(10, 20))
	fmt.Println(c.div(10, 20))
}

