package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	// func greeterTwo() {              // YOU CAN'T DEFINE A FUNCTION INSIDE ANOTHER FUNCTION
	// 	fmt.Println("Another method")
	// }
	// greeterTwo()


	// Adding Two numbers
	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	// Adding All values
	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", myMessage)
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) { // you need to provide values and all of the values are of type integer
	total := 0
	for _, val := range values{
		total += val
	}
	return total, "Hi Pro result function"
}