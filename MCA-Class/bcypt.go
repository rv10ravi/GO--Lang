
//Note:  go mod init bcypt.go and  go get -u golang.org/x/crypto/bcrypt

package main

import (
"fmt" // Importing the fmt package for formatted I/O
"golang.org/x/crypto/bcrypt" // Importing the bcrypt package for password hashing
)

func main() {
// Example password
password := "examplePassword123"

// Hashing the password
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// GenerateFromPassword generates a hashed password from the given plaintext password.
// It takes the plaintext password as a byte slice and the cost parameter (which determine hashing cost and thus the computational effort required).
// It returns the hashed password and an error, if any.

if err != nil { // Check if there&#39;s an error while generating the hashed password
fmt.Println("Error hashing password:", err) // Print the error message
return // Return from the function
}

fmt.Println("Hashed password:", string(hashedPassword)) // Print the hashed password

// Comparing hashed password with plaintext password
plaintextPassword := "examplePassword123" // Assigning the plaintext password for comparison

// CompareHashAndPassword compares a hashed password with its possible plaintext equivalent.
// It takes the hashed password and the plaintext password as byte slices.
// It returns nil on success or an error if the comparison fails.
err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(plaintextPassword))

if err != nil { // Check if there&#39;s an error during comparison
fmt.Println("Password does not match hash:", err) // Print the error message
return // Return from the function
}

fmt.Println("Password matches hash.") // Print success message
}