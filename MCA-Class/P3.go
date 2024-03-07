package main

import (
    "errors"
    "fmt"
    "time"
)

type ProductType int

const (
    Electronics ProductType = iota
    Clothing
    Groceries
)

type Product struct {
    ID          int
    Name        string
    Price       float64
    Type        ProductType
    Description string
}

type ElectronicsProduct struct {
    Product
    WarrantyPeriod time.Duration
}

type ClothingProduct struct {
    Product
    Size string
}

type GroceryProduct struct {
    Product
    ExpirationDate time.Time
}

// Storage defines the interface for different storage options
type Storage interface {
    SaveProduct(product Product) error
    UpdateProduct(product Product) error
    DeleteProduct(productID int) error
    GetProduct(productID int) (Product, error)
}

// FileStorage represents local file storage
type FileStorage struct {
    products map[int]Product
}

func NewFileStorage() *FileStorage {
    return &FileStorage{
        products: make(map[int]Product),
    }
}

func (fs *FileStorage) SaveProduct(product Product) error {
    fs.products[product.ID] = product
    return nil
}

func (fs *FileStorage) UpdateProduct(product Product) error {
    if _, ok := fs.products[product.ID]; !ok {
        return errors.New("product not found")
    }
    fs.products[product.ID] = product
    return nil
}

func (fs *FileStorage) DeleteProduct(productID int) error {
    if _, ok := fs.products[productID]; !ok {
        return errors.New("product not found")
    }
    delete(fs.products, productID)
    return nil
}

func (fs *FileStorage) GetProduct(productID int) (Product, error) {
    product, ok := fs.products[productID]
    if !ok {
        return Product{}, errors.New("product not found")
    }
    return product, nil
}

func main() {
    // Main function for testing the implementation
    fmt.Println("Inventory Management System")

    // Initialize a new file storage
    fileStorage := NewFileStorage()

    // Test adding a new product
    product1 := Product{ID: 1, Name: "Laptop", Price: 1000.0, Type: Electronics, Description: "High-performance laptop"}
    fileStorage.SaveProduct(product1)

    // Test retrieving the added product
    retrievedProduct, err := fileStorage.GetProduct(1)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Retrieved Product:", retrievedProduct)
    }

    // Test updating the product
    updatedProduct := Product{ID: 1, Name: "Laptop Pro", Price: 1200.0, Type: Electronics, Description: "Upgraded laptop"}
    fileStorage.UpdateProduct(updatedProduct)

    // Test retrieving the updated product
    retrievedUpdatedProduct, err := fileStorage.GetProduct(1)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Retrieved Updated Product:", retrievedUpdatedProduct)
    }

    // Test deleting the product
    err = fileStorage.DeleteProduct(1)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Product deleted successfully")
    }

    // Test retrieving the deleted product
    _, err = fileStorage.GetProduct(1)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
