package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in no golang; No super or parent concept exists in golang;

	diptansh := User{"Vanshika", "vanshika@go.dev", true, 16}
	fmt.Println(diptansh)
	fmt.Printf("Diptansh details are: %+v\n", diptansh)
	fmt.Printf("Name is %v and email is %v", diptansh.Name, diptansh.Email)
}

type User struct {
	Name	string
	Email	string
	Status	bool
	Age		int
}