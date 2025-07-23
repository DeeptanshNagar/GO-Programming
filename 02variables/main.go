package main

import "fmt"

func main() {
	var username string = "deeptansh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255 // 256 will give you an error bcoz range is 0-255
	// uint16 -> 0 to 65535 (0 to 2¹⁶ - 1)
	// uint 32 -> 0 to 4,294,967,295 (0 to 2³² - 1)
	// unit 64 -> 0 to 18,446,744,073,709,551,615 (0 to 2⁶⁴ - 1)
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)
}
