package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Bookstore struct {
	stock map[string]int 
	mu    sync.Mutex     // Mutex to synchronize access to stock
}


func NewBookstore() *Bookstore {
	return &Bookstore{
		stock: map[string]int{
			"Harry Potter":  1,
			"Lord of Rings": 15,
			"Game of Thrones": 20,
			"Atomic Habits": 5,
		},
	}
}


func (b *Bookstore) PurchaseBook(book string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	b.mu.Lock()
	defer b.mu.Unlock()

	fmt.Printf("Customer is trying to purchase %s\n", book)

	if stock, ok := b.stock[book]; ok && stock > 0 {
		// Simulating purchase delay
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		b.stock[book]--
		ch <- fmt.Sprintf("Purchased %s", book)
		fmt.Printf("Customer purchased %s\n", book)
	} else {
		ch <- fmt.Sprintf("Book %s is out of stock", book)
		fmt.Printf("Book %s is out of stock\n", book)
	}
}

func main() {
	bookstore := NewBookstore()

	// Create a channel to communicate purchase statuses
	purchaseCh := make(chan string)

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Simulate multiple customers purchasing books concurrently
	for _, book := range []string{"Harry Potter", "Lord of Rings", "Game of Thrones", "Atomic Habits", "Harry Potter"} {
		wg.Add(1)
		go bookstore.PurchaseBook(book, &wg, purchaseCh)
	}

	// Close the channel after all purchases are done
	go func() {
		wg.Wait()
		close(purchaseCh)
	}()

	// Print purchase statuses
	for purchase := range purchaseCh {
		fmt.Println(purchase)
	}
}
