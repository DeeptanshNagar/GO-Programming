package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// Note: Go does not support inheritance like traditional OOP languages.
	// There are no concepts like super classes or parent classes in Go.

	// Creating an instance of the User struct
	diptansh := User{"Vanshika", "vanshika@go.dev", true, 16}

	// Printing the struct directly
	fmt.Println(diptansh)

	// Printing struct fields using formatted output
	fmt.Printf("Diptansh details are: %+v\n", diptansh)
	fmt.Printf("Name is %v and email is %v\n", diptansh.Name, diptansh.Email)

	// Calling a method with a value receiver (doesn't modify the original struct)
	diptansh.GetStatus()

	// Calling a method that attempts to change the email using a value receiver.
	// This won't affect the original struct because it's working on a copy.
	diptansh.NewMail()

	// Verifying that the original struct remains unchanged after calling NewMail()
	fmt.Printf("Name is %v and email is %v\n", diptansh.Name, diptansh.Email)
}

// User represents a user in the system
type User struct {
	Name   string // Exported (public) because it starts with an uppercase letter
	Email  string // Exported (public)
	Status bool   // Exported (public)
	Age    int    // Exported (public)

	// oneAge int // Unexported (private) because it starts with a lowercase letter
}

// GetStatus prints whether the user is active (value receiver)
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

// NewMail tries to change the user's email but won't succeed
// in modifying the original struct (uses value receiver)
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email inside NewMail():", u.Email)
}
