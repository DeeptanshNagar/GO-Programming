package main

import (
	"bufio"
	"fmt"
	"os"
)



func main () {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n') // same as try catch concept of java means either u will get input or you'll receive an error 
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Thanks for rating is %T", input)
}