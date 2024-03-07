// //Implement the concepts of Interfaces using basic calculator

// package main

// import "fmt"

// type calculator interface {
// 	add(a int, b int) int
// 	sub(a int, b int) int
// 	mul(a int, b int) int
// 	div(a int, b int) int
// }

// type calc struct {
// }

// func (c calc) add(a int, b int) int {
// 	return a + b
// }

// func (c calc) sub(a int, b int) int {
// 	return a - b
// }

// func (c calc) mul(a int, b int) int {
// 	return a * b
// }

// func (c calc) div(a int, b int) int {
// 	return a / b
// }

// func main() {
// 	var c calculator
// 	c = calc{}
// 	fmt.Println(c.add(10, 20))
// 	fmt.Println(c.sub(10, 20))
// 	fmt.Println(c.mul(10, 20))
// 	fmt.Println(c.div(10, 20))
// }

package main

import "fmt"

type book struct {
	isbn int64
	title string
	author string
	price float64
}

type bookdetails interface {
	getBookIsbn() int64
	getBookTitle() string
	getBookAuthor() string
	getBookPrice() float64
}

func (b book) getBookIsbn() int64 {
	return b.isbn
}

func (b book) getBookTitle() string {
	return b.title
}

func (b book) getBookAuthor() string {
	return b.author
}

func (b book) getBookPrice() float64 {
	return b.price
}

func main() {
	b := book{
		isbn:  12345,
		title: "Go Programming",
		author: "Ravi",
		price:  500,
	}
	fmt.Println(b.getBookIsbn())
	fmt.Println(b.getBookTitle())
	fmt.Println(b.getBookAuthor())
	fmt.Println(b.getBookPrice())
}


