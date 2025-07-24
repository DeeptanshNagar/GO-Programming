package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	// Syntax 1
	loginCount := 9
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	// Syntax 2
	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	// Syntax 3
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	// Syntax 4
	// if err != nil {
		
	// }
}